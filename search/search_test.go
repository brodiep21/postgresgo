package search_test

import (
	"testing"

	"github.com/brodiep21/postgresgo/search"
)

func TestSearchHorsepower(t *testing.T) {
	car := "infiniti qx80"

	want := "400"
	got, _ := search.DataSearch(car)

	if got != want {
		t.Errorf("Got %s, want %s", got, want)
	}

}
func TestSearchMSRP(t *testing.T) {
	car := "infiniti qx80"

	want := "71,100"
	_, got := search.DataSearch(car)

	if got != want {
		t.Errorf("Got %s, want %s", got, want)
	}

}
