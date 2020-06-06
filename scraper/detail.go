package scraper

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func FetchDetail(url string) (string, string, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("ERROR: can't get response, %#v\n", err.Error())
		return "", "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Printf("status code error: %d %s", resp.StatusCode, resp.Status)
		return "", "", err
	}

	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Printf("ERROR: failed to fetch html, %#v\n", err.Error())
		panic(err)
	}

	author := doc.Find(".author").Text()
	if author == "" {
		author = "不明"
	}

	paragraph := doc.Find("p").Text()

	return author, string([]rune(paragraph)[:50]), err
}
