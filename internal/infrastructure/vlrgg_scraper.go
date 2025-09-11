package infrastructure

import (
	"fmt"
	"log/slog"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/miztch/valorant-match-schedule/internal/domain"
)

const (
	vlrGGDomain             = "www.vlr.gg"
	baseURL                 = "https://" + vlrGGDomain
	crawlerUserAgent        = "Googlebot/2.1 (+http://www.google.com/bot.html)"
	minDelayInSeconds       = 3
	maxRandomDelayInSeconds = 7
)

// VlrGGScraper is a scraper for vlr.gg
type VlrGGScraper struct {
	Collector *colly.Collector
	BaseURL   string
}

// NewVlrGGScraper creates a new VlrGGScraper
func NewVlrGGScraper() *VlrGGScraper {
	return &VlrGGScraper{
		Collector: SetupColly(),
		BaseURL:   baseURL,
	}
}

// SetupColly sets up a colly collector
func SetupColly() *colly.Collector {
	c := colly.NewCollector(
		colly.AllowedDomains(vlrGGDomain),
		colly.UserAgent(crawlerUserAgent),
	)

	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Delay:       minDelayInSeconds * time.Second,
		RandomDelay: maxRandomDelayInSeconds * time.Second,
	},
	)

	c.OnError(func(r *colly.Response, err error) {
		slog.Error("Request failed", "url", r.Request.URL.String(), "error", err.Error(), "statusCode", r.StatusCode)
	})

	return c
}

// buildRequestURL builds a request URL for vlr.gg
func (v *VlrGGScraper) buildRequestURL(path string) string {
	return v.BaseURL + path
}

// getMatchURLList gets a list of match URLs from vlr.gg
func (v *VlrGGScraper) getMatchURLList(pageNumber int) ([]string, error) {
	requestURL := v.buildRequestURL("/matches?page=" + strconv.Itoa(pageNumber))
	slog.Info("Scraping match URLs", "url", requestURL, "page", pageNumber)

	var matchURLList []string
	v.Collector.OnHTML(".match-item", func(e *colly.HTMLElement) {
		matchUrlPath := e.Attr("href")
		matchURLList = append(matchURLList, matchUrlPath)
	})

	err := v.Collector.Visit(requestURL)

	if err != nil {
		slog.Error("Failed to scrape match URLs", "error", err.Error(), "url", requestURL)
		return nil, fmt.Errorf("failed to visit %s: %w", requestURL, err)
	}

	slog.Info("Successfully scraped match URLs", "count", len(matchURLList), "page", pageNumber)
	return matchURLList, nil
}

// parseScrapedEvent parses a scraped event from vlr.gg
func parseScrapedEvent(e *colly.HTMLElement, eventUrlPath string) (domain.VlrEvent, error) {
	idStr := strings.Split(eventUrlPath, "/")[2]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return domain.VlrEvent{}, err
	}

	vlrEvent := domain.VlrEvent{
		Id:   id,
		Name: e.ChildText(".wf-title"),
	}

	countryFlagPlaceHolder := e.ChildAttr(".event-desc-item-value > .flag", "class")
	r := strings.NewReplacer("flag mod-", "")
	countryFlag := r.Replace(countryFlagPlaceHolder)

	vlrEvent.CountryFlag = countryFlag
	return vlrEvent, nil
}

// scrapeEvent scrapes an event from vlr.gg
func (v *VlrGGScraper) scrapeEvent(eventUrlPath string) (domain.VlrEvent, error) {
	requestURL := v.buildRequestURL(eventUrlPath)

	var event domain.VlrEvent
	v.Collector.OnHTML(".event-header", func(e *colly.HTMLElement) {
		var err error
		event, err = parseScrapedEvent(e, eventUrlPath)
		if err != nil {
			slog.Error("Failed to parse scraped event", "error", err.Error(), "url", eventUrlPath)
			return
		}
	})

	err := v.Collector.Visit(requestURL)

	if err != nil {
		return domain.VlrEvent{}, fmt.Errorf("failed to visit %s: %w", requestURL, err)
	}

	return event, nil
}

// parseScrapedMatch parses a scraped match from vlr.gg
func parseScrapedMatch(e *colly.HTMLElement, matchUrlPath string) domain.VlrMatch {
	matchId, _ := strconv.Atoi(strings.Split(matchUrlPath, "/")[1])

	vlrMatch := domain.VlrMatch{
		Id:            matchId,
		PagePath:      matchUrlPath,
		EventPagePath: e.ChildAttr(".match-header-event", "href"),
	}

	// Name
	matchNamePlaceHolder := e.ChildText(".match-header-event-series")
	r := strings.NewReplacer("\t", "", "\n", "")
	matchName := r.Replace(matchNamePlaceHolder)
	vlrMatch.Name = matchName

	// StartTime, StartDate
	// convert EST -> UTC
	startTimePlaceHolder := e.ChildAttr(".moment-tz-convert", "data-utc-ts")
	loc, _ := time.LoadLocation("America/New_York")
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", startTimePlaceHolder, loc)
	utc := t.UTC()
	vlrMatch.StartTime = utc.Format("2006-01-02T15:04:05+0000")
	vlrMatch.StartDate = utc.Format("2006-01-02")

	// BestOf
	bestOfPlaceHolder := e.ChildText(".match-header-vs-score > .match-header-vs-note:last-of-type")
	r = strings.NewReplacer("Bo", "", " Maps", "", "\t", "", "\n", "")
	bestOf, err := strconv.Atoi(r.Replace(bestOfPlaceHolder))
	if err != nil {
		return domain.VlrMatch{}
	}
	vlrMatch.BestOf = bestOf

	// Teams
	teams := []domain.Team{}
	e.ForEach(".wf-title-med", func(_ int, el *colly.HTMLElement) {
		var t domain.Team
		teamNamePlaceHolder := el.Text
		r = strings.NewReplacer("\t", "", "\n", "")
		t.Name = r.Replace(teamNamePlaceHolder)
		teams = append(teams, t)
	})
	vlrMatch.Teams = teams

	return vlrMatch
}

// ScrapeMatch scrapes a match from vlr.gg
func (v *VlrGGScraper) scrapeMatch(matchUrlPath string) (domain.VlrMatch, error) {
	requestURL := v.buildRequestURL(matchUrlPath)
	slog.Info("Scraping match details", "url", requestURL, "matchPath", matchUrlPath)

	var match domain.VlrMatch
	v.Collector.OnHTML(".match-header", func(e *colly.HTMLElement) {
		match = parseScrapedMatch(e, matchUrlPath)
	})

	err := v.Collector.Visit(requestURL)

	if err != nil {
		slog.Error("Failed to scrape match details", "error", err.Error(), "url", requestURL)
		return domain.VlrMatch{}, fmt.Errorf("failed to visit %s: %w", requestURL, err)
	}

	slog.Info("Successfully scraped match",
		"matchId", match.Id,
		"matchName", match.Name,
		"startTime", match.StartTime,
		"bestOf", match.BestOf,
	)
	return match, nil
}
