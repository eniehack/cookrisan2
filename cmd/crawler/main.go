package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	target_url := flag.String("url", "", "")
	flag.Parse()
	if len(*target_url) < 1 {
		os.Exit(1)
	}
	client := new(http.Client)
	resp, err := client.Get(*target_url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("bad status code returned: %s\n", resp.Status)
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	doc.Find(".recipe-detail-ingredient__name").Each(func(i int, s *goquery.Selection) {
		nolink := s.Find(".recipe-detail-ingredient__name__type--nolink").Text()
		if nolink != "" {
			log.Printf("NoLink Ingredient: %s\n", nolink)
		}

		// a .recipe-detail-ingredient__name__type のテキストを取得
		link := s.Find(".recipe-detail-ingredient__name__type").Text()
		if link != "" {
			log.Printf("Link Ingredient: %s\n", link)
		}
	})
}
