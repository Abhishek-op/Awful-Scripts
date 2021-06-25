package main

import (
  "fmt"
	"github.com/gocolly/colly"
)

func main() {
	give_current_price_of_gold()
}

func give_current_price_of_gold() {
	c := colly.NewCollector(
		colly.AllowedDomains("google.com", "www.google.com"),
	)

	//bool variable to close the search after target tag is found
	found := false

	c.OnHTML("div", func(e *colly.HTMLElement) {

		if found {
			return
		}

		if e.Attr("class") == "BNeawe s3v9rd AP7Wnd" {
			fmt.Printf("1 BTC = %s\n", e.Text)
			found = true
		}
	})

	//webpage url to be visited
	c.Visit("https://www.google.com/search?q=gold+price")
}