package search

import (
	"context"
	"fmt"
	"log"
	"strings"

	googlesearch "github.com/rocketlaunchr/google-search"
)

func HorsepowerSearch(vehicle string) string {

	var hp = "horsepower"
	var carHP string
	res, err := googlesearch.Search(context.TODO(), vehicle+" "+hp)
	if err != nil {
		log.Fatalf("Could not get horsepower information, %s", err)
	}

	for _, v := range res {

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

	return carHP

}

func MsrpSearch(vehicle string) string {

	var msrp = "msrp"
	var carmsrp string
	res, err := googlesearch.Search(context.TODO(), "infiniti qx80"+" "+msrp)
	if err != nil {
		log.Fatalf("Could not get search information, %s", err)
	}

	for _, v := range res {

		if strings.Contains(v.Description, "$") {
			pos := strings.Index(v.Description, "$")
			if pos == -1 {
				log.Fatal("Could not gather MSRP information, please try to refine the search to an exact Make and Model, you are welcome to include year")
			}
			carmsrp = v.Description[pos : pos+7]
			carmsrp = strings.ReplaceAll(carmsrp, "$", "")
			fmt.Println(carmsrp)
			break
		} else {
			continue
		}
	}
	return carmsrp
}
