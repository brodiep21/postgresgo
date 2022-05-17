package search

import (
	"context"
	"log"
	"strings"

	googlesearch "github.com/rocketlaunchr/google-search"
)

//returns vehicle's horsepower according to the first URL google search
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
				log.Fatal("Could not gather MSRP information, please try to refine the search to an exact Make and Model, you are welcome to include year")
			}
			carHP = v.Description[(pos - 4):pos]
			break
		} else {
			continue
		}
	}
	//cut out any spacing or dashes
	carHP = strings.ReplaceAll(carHP, " ", "")
	carHP = strings.ReplaceAll(carHP, "-", "")
	return carHP

}

//returns vehicle's MSRP according to the first URL google search
func MsrpSearch(vehicle string) string {

	var msrp = "msrp"
	var carmsrp string
	res, err := googlesearch.Search(context.TODO(), vehicle+" "+msrp)
	if err != nil {
		log.Fatalf("Could not get MSRP information, %s", err)
	}

	for _, v := range res {

		if strings.Contains(v.Description, "$") {
			pos := strings.Index(v.Description, "$")
			if pos == -1 {
				log.Fatal("Could not gather MSRP information, please try to refine the search to an exact Make and Model, you are welcome to include year")
			}
			carmsrp = v.Description[pos : pos+8]
			carmsrp = strings.ReplaceAll(carmsrp, "$", "")
			break
		} else {
			continue
		}
	}
	//filters out any stray spaces or periods, commas at the end of the string.
	carmsrp = strings.ReplaceAll(carmsrp, " ", "")
	carmsrp = strings.ReplaceAll(carmsrp, ".", "")
	carmsrp = strings.TrimSuffix(carmsrp, ",")
	return carmsrp
}
