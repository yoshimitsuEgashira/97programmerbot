package scraper

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func FetchEssay(n int, url string) (string, string, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("ERROR: can't get response, %#v\n", err.Error())
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

	var title, path string

	doc.Find("ol li").Each(func(i int, s *goquery.Selection) {
		if i == n {
			title = s.Find("a").Text()
			path, _ = s.Find("a").Attr("href")
		}
	})
	return title, url + path, nil
}
