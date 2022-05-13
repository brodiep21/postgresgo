package scraper

import (
	"context"
	"fmt"
	"log"
	"strings"

	googlesearch "github.com/rocketlaunchr/google-search"
)

func Webscraper(search string) (string, string) {

	res, err := googlesearch.Search(context.TODO(), search)
	if err != nil {
		log.Fatalf("Could not get search information, %s", err)
	}

	// var descripts string
	var horsepower = "hp"
	for _, v := range res {
		// fmt.Println(v.Description)
		if strings.Contains(v.Description, horsepower) {
			pos := strings.Index(v.Description, horsepower)
			if pos == -1 {
				return
			}
			horsepower = v.Description[(pos - 5):pos]

			fmt.Println(horsepower)

		} else {
			continue
		}
	}
	// c := colly.NewCollector(
	// 	colly.AllowedDomains("www.google.com"),
	// )

	// var msrpCost string
	// //grab data for msrp base
	// c.OnHTML("kc:/automotive/model_year:min msrp", func(e *colly.HTMLElement) {
	// 	subDiv := e.ChildText("span.LrzXr.kno-fv.wHYlTd.z8gr9e")
	// 	subDiv = strings.ToLower(subDiv)
	// 	msrpCost = strings.ReplaceAll(subDiv, "from $", "")
	// })

	// var horsepower string
	// //grab data for horsepower
	// c.OnHTML("kc:/automotive/model_year:horsepower", func(e *colly.HTMLElement) {
	// 	subDiv := e.ChildText("span.LrzXr.kno-fv.wHYlTd.z8gr9e")
	// 	horsepower = strings.ReplaceAll(subDiv, " hp", "")
	// })

	// c.OnRequest(func(request *colly.Request) {
	// 	fmt.Println("Visiting", request.URL.String())
	// })

	// search = strings.ReplaceAll(search, " ", "+")
	// c.Visit("https://www.google.com/search?q=" + search)

	// fmt.Println(msrpCost, horsepower)

	// return msrpCost, horsepower

}
