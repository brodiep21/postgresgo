package search

import (
	"context"
	"fmt"
	"log"
	"strings"

	googlesearch "github.com/rocketlaunchr/google-search"
)

func DataSearch(search string) (string, string) {

	var hp = "horsepower"
	var carHP string
	res, err := googlesearch.Search(context.TODO(), search+" "+hp)
	if err != nil {
		log.Fatalf("Could not get search information, %s", err)
	}

	for _, v := range res {
		// fmt.Println(v.Description)
		if strings.Contains(v.Description, hp) {
			pos := strings.Index(v.Description, hp)
			if pos == -1 {
				break
			}
			carHP = v.Description[(pos - 4):pos]
			break
		} else {
			continue
		}
	}

	var msrp = "msrp"
	var carmsrp string
	res, err = googlesearch.Search(context.TODO(), search+" "+msrp)
	if err != nil {
		log.Fatalf("Could not get search information, %s", err)
	}

	for _, v := range res {
		fmt.Println(v.Description)
		if strings.Contains(v.Description, "starting MSRP") {
			pos := strings.Index(v.Description, "starting MSRP")
			if pos == -1 {
				log.Fatal("Could not gather MSRP information, please try to refine the search to an exact Make and Model, you are welcome to include year")
			}
			carmsrp = v.Description[pos+13 : pos+19]
			fmt.Println(carmsrp)
			break
		} else {
			continue
		}
	}

	return carHP, msrp

}
