package search

import (
	"context"
	"errors"
	"strings"

	googlesearch "github.com/rocketlaunchr/google-search"
)

//returns vehicle's horsepower according to the first URL google search that has corresponding Horsepower information
func HorsepowerSearch(vehicle string) (string, error) {

	var hp = "horsepower"
	var carHP string
	res, err := googlesearch.Search(context.TODO(), vehicle+" "+hp)
	if err != nil {
		return "", err
	}
	for _, v := range res {

		if strings.Contains(v.Description, hp) {
			pos := strings.Index(v.Description, hp)
			if pos == -1 {
				return "", errors.New("Could not find the Horsepower of " + vehicle)
			}
			carHP = v.Description[(pos - 4):pos]
			break
		}
	}
	//cut out any spacing or dashes
	carHP = strings.ReplaceAll(carHP, " ", "")
	carHP = strings.ReplaceAll(carHP, "-", "")
	return carHP, nil

}

//returns vehicle's MSRP according to the first URL google search that has corresponding $ information
func MsrpSearch(vehicle string) (string, error) {

	var msrp = "msrp"
	var carmsrp string
	res, err := googlesearch.Search(context.TODO(), vehicle+" "+msrp)
	if err != nil {
		return "", err
	}

	for _, v := range res {

		if strings.Contains(v.Description, "$") {
			pos := strings.Index(v.Description, "$")
			if pos == -1 {
				return "", errors.New("Could not find the MSRP of " + vehicle)
			}
			carmsrp = v.Description[pos : pos+8]
			carmsrp = strings.ReplaceAll(carmsrp, "$", "")
			break
		}
	}
	//filters out any stray spaces or periods, commas at the end of the string.
	carmsrp = strings.ReplaceAll(carmsrp, " ", "")
	carmsrp = strings.ReplaceAll(carmsrp, ".", "")
	carmsrp = strings.TrimSuffix(carmsrp, ",")
	return carmsrp, nil
}
