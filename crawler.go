package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(
	// colly.AllowedDomains("aircraftcompare.com/"),
	)

	// Callback for when a scraped page contains a article element
	c.OnHTML("article", func(e *colly.HTMLElement) {
		isAirplanePage := false

		// Extract div tags from the document
		e.DOM.ParentsUntil("~").Find("meta").Each(func(_ int, s *goquery.Selection) {
			// Search for section type tags
			if property, _ := s.Attr("property"); strings.EqualFold(property, "og:type") {
				content, _ := s.Attr("content")

				// Airplane pages have article as their class
				isAirplanePage = strings.EqualFold(content, "article")
			}
		})

		if isAirplanePage {
			// Find the Airplane page hrefs
			fmt.Println("Airplane: ", e.DOM.Find("h1").Text())
			// Grab all of the descriptions of the airplanes
			fmt.Println("Description: ", e.DOM.Find("p").Text())
		}
	})

	// Callback for links on scraped pages
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		// Extract the linked URL from the anchor tag
		link := e.Attr("href")
		// Have our crawler visit the linked URL
		c.Visit(e.Request.AbsoluteURL(link))
	})

	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		RandomDelay: 1 * time.Second,
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://www.aircraftcompare.com/")
}
