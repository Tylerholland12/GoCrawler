package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type info struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
	Link string `json:"link"`
}

func main() {
	c := colly.NewCollector()

	// planeInfo := new(info)
	// planeInfo.Name = ""
	// planeInfo.Desc = ""
	// planeInfo.Link = ""

	c.OnHTML("section[class='row row-small-gutter']", func(e *colly.HTMLElement) {
		link := e.Attr("class")

		fmt.Println("first step:", "-->", link)
	})

	c.OnHTML("header[class]", func(e *colly.HTMLElement) {
		link := e.Attr("class")

		fmt.Println("second step:", "-->", link)
	})

	c.OnHTML("a[href='https://www.aircraftcompare.com/aircraft-categories/commercial-airplanes/']", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		fmt.Println("third step:", "-->", link)
	})

	c.Visit("https://www.aircraftcompare.com/")
}
