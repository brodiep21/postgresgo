package scraper

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func Webscraper(search string) (string, string) {

	c := colly.NewCollector(
		colly.AllowedDomains("www.google.com"),
	)

	var msrpCost string
	//grab data for msrp base
	c.OnHTML("kc:/automotive/model_year:min msrp", func(e *colly.HTMLElement) {
		subDiv := e.ChildText("span.LrzXr.kno-fv.wHYlTd.z8gr9e")
		subDiv = strings.ToLower(subDiv)
		msrpCost = strings.ReplaceAll(subDiv, "from $", "")
	})

	var horsepower string
	//grab data for horsepower
	c.OnHTML("kc:/automotive/model_year:horsepower", func(e *colly.HTMLElement) {
		subDiv := e.ChildText("span.LrzXr.kno-fv.wHYlTd.z8gr9e")
		horsepower = strings.ReplaceAll(subDiv, " hp", "")
	})

	c.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String())
	})

	search = strings.ReplaceAll(search, " ", "+")
	c.Visit("https://www.google.com/search?q=" + search)

	fmt.Println(msrpCost, horsepower)

	return msrpCost, horsepower

}
