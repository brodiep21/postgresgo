package scraper_test

import (
	"testing"

	"github.com/brodiep21/postgresgo/scraper"
)

func TestWebScraperFirstResponse(t *testing.T) {
	car := "infiniti qx80"

	want := "71,100"
	got, _ := scraper.Webscraper(car)

	if got != want {
		t.Errorf("Got %s, want %s", got, want)
	}

}
func TestWebScraperSecondResponse(t *testing.T) {
	car := "infiniti qx80"

	want := "400"
	_, got := scraper.Webscraper(car)

	if got != want {
		t.Errorf("Got %s, want %s", got, want)
	}

}
