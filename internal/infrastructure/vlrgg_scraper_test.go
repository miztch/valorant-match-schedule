package infrastructure

import (
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"
)

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
