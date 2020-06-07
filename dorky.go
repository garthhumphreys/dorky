package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://go-colly.org/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("Response status", resp.Status)

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("#menu a").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})

}