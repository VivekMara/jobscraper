package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main(){
	c := colly.NewCollector(
	)

	c.OnHTML("a[href]", func(h *colly.HTMLElement) {
		link := h.Attr("href")
		fmt.Printf("Link found: %q -> %s", h.Text, link)
		c.Visit(h.Request.AbsoluteURL(link))
	})
	
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://wikipedia.com")
}