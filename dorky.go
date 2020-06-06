package main

import (
	"fmt"
	"github.com/gocolly/colly"
	// "os"
)

func main() {
	// sitePtr := flag.String("site", "gov.ky", "a domain")
	// inurlPtr := flag.String("inurl", "admin", "what to look for in the url")
	// wildcardPtr := flag.Bool("w", true, "search sub-domains")
	// resultsPtr := flag.Int("r", 10, "number of dork results to show")
	//
	// flag.Parse()
	// fmt.Println("site:", *sitePtr)
	// fmt.Println("inurl:", *inurlPtr)
	// fmt.Println("w:", *wildcardPtr)
	// fmt.Println("r:", *resultsPtr)
	// fmt.Println("other params:", flag.Args())

	/*
	Use this for just arguments loose form, instead of flags
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]
	arg := os.Args[3]
	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
	*/

	search()
}

func search()  {
	fmt.Println("Search")
	c := colly.NewCollector()
	c.OnHTML("#menu .computer a", func(element *colly.HTMLElement) {
		fmt.Printf("title:%s link:%s\n", element.Text, element.Attr("href"))
	})

	c.OnRequest(func(request *colly.Request) {
		fmt.Println("Visting", request.URL)
	})

	c.Visit("http://go-colly.org/")
}