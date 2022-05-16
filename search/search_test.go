package search_test

import (
	"testing"

	"github.com/brodiep21/postgresgo/search"
)

func TestSearchHorsepower(t *testing.T) {
	car := "infiniti qx80"

	want := "400"
	got := search.HorsepowerSearch(car)

	if got != want {
		t.Errorf("Got %s, want %s", got, want)
	}

}
func TestMSRPsearch(t *testing.T) {
	car := "infiniti qx80"

	want := "71,100"
	got := search.MsrpSearch(car)

	if got != want {
		t.Errorf("Got %s, want %s", got, want)
	}

}
