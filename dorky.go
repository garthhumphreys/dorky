package main

import (
	"flag"
	"fmt"
	"github.com/gocolly/colly"
	"os"

	// "os"
)

func main() {
	sitePtr := flag.String("site", "gov.ky", "a domain")
	inurlPtr := flag.String("inurl", "admin", "what to look for in the url")
	wildcardPtr := flag.Bool("w", true, "search sub-domains")
	resultsPtr := flag.Int("r", 10, "number of dork results to show")

	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Println("expected 'site'")
		os.Exit(1)
	}

	if *sitePtr != "" {
		//search(*sitePtr)
		fmt.Println("site:", *sitePtr)
		search("https://www.google.com/search?q=site%3A.ky+inurl%3Aphp")
	} else {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *inurlPtr != "" {
		fmt.Println("inurl:", *inurlPtr)
	}

	if *wildcardPtr {
		fmt.Println("w:", *wildcardPtr)
	}

	if *resultsPtr == 10 {
		fmt.Println("r:", *resultsPtr)
	}

	/*
	Use this for just arguments loose form, instead of flags
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]
	arg := os.Args[3]
	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
	*/

	//search("http://go-colly.org/")
}

func search(url string)  {
	fmt.Println("Search")
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/47.0"),
		)
	c.OnHTML("#search .r > a", func(element *colly.HTMLElement) {
		fmt.Printf("title:%s link:%s\n", element.Text, element.Attr("href"))
	})

	c.OnRequest(func(request *colly.Request) {
		fmt.Println("Visting", request.URL)
	})

	c.Visit(url)
}