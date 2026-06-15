package infrastructure

import (
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"

	"github.com/gocolly/colly/v2"
)

func serveHTML(t *testing.T, html string) (url string, close func()) {
	t.Helper()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(html))
	}))
	return ts.URL, ts.Close
}

func TestParseScrapedEvent(t *testing.T) {
	tests := []struct {
		name      string
		html      string
		eventPath string
		wantId    int
		wantName  string
		wantFlag  string
	}{
		{
			name: "new layout: event-header-main-title",
			html: `<html><body><div class="event-header">
				<h1 class="event-header-main-title">VCT 2026: Pacific Stage 2</h1>
				<div class="event-desc-items">
					<div class="event-desc-item-value"><i class="flag mod-jp"></i>Japan</div>
				</div>
			</div></body></html>`,
			eventPath: "/event/2776/vct-2026-pacific-stage-2",
			wantId:    2776,
			wantName:  "VCT 2026: Pacific Stage 2",
			wantFlag:  "jp",
		},
		{
			name: "old layout: wf-title with direct child flag",
			html: `<html><body><div class="event-header">
				<h1 class="wf-title">VCT 2025: Pacific Stage 1</h1>
				<div class="event-desc-item-value"><i class="flag mod-kr"></i>Korea</div>
			</div></body></html>`,
			eventPath: "/event/1234/vct-2025-pacific-stage-1",
			wantId:    1234,
			wantName:  "VCT 2025: Pacific Stage 1",
			wantFlag:  "kr",
		},
		{
			name: "no flag (global/online event)",
			html: `<html><body><div class="event-header">
				<h1 class="event-header-main-title">Challengers 2026: Japan Season Finals</h1>
			</div></body></html>`,
			eventPath: "/event/2991/challengers-2026-japan-season-finals",
			wantId:    2991,
			wantName:  "Challengers 2026: Japan Season Finals",
			wantFlag:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url, close := serveHTML(t, tt.html)
			defer close()

			c := colly.NewCollector()
			var id int
			var name, flag string
			c.OnHTML(".event-header", func(e *colly.HTMLElement) {
				result, err := parseScrapedEvent(e, tt.eventPath)
				if err != nil {
					t.Errorf("unexpected error: %v", err)
					return
				}
				id = result.Id
				name = result.Name
				flag = result.CountryFlag
			})
			c.Visit(url)

			if id != tt.wantId {
				t.Errorf("Id = %d, want %d", id, tt.wantId)
			}
			if name != tt.wantName {
				t.Errorf("Name = %q, want %q", name, tt.wantName)
			}
			if flag != tt.wantFlag {
				t.Errorf("CountryFlag = %q, want %q", flag, tt.wantFlag)
			}
		})
	}
}

func TestRetryTransport(t *testing.T) {
	tests := []struct {
		name             string
		failCount        int
		wantRequestCount int
		wantStatusCode   int
	}{
		{
			name:             "succeeds after retries within maxRetries",
			failCount:        3,
			wantRequestCount: 4, // 3 failures + 1 success
			wantStatusCode:   http.StatusOK,
		},
		{
			name:             "fails after exhausting maxRetries",
			failCount:        maxRetries + 1,
			wantRequestCount: maxRetries + 1,
			wantStatusCode:   http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var requestCount atomic.Int32
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				count := int(requestCount.Add(1))
				if count <= tt.failCount {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				w.WriteHeader(http.StatusOK)
			}))
			defer ts.Close()

			transport := &retryTransport{
				base:             http.DefaultTransport,
				maxRetries:       maxRetries,
				initialRetryWait: 1 * time.Millisecond,
			}
			client := &http.Client{Transport: transport}

			resp, err := client.Get(ts.URL)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			defer resp.Body.Close()

			if got := int(requestCount.Load()); got != tt.wantRequestCount {
				t.Errorf("request count = %d, want %d", got, tt.wantRequestCount)
			}
			if resp.StatusCode != tt.wantStatusCode {
				t.Errorf("status code = %d, want %d", resp.StatusCode, tt.wantStatusCode)
			}
		})
	}
}
