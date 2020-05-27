package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const (
	url string = "https://xn--97-273ae6a4irb6e2hsoiozc2g4b8082p.com"
	max int    = 106
)

func GetEssay(n int) (string, string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}
	fmt.Printf("status: %d\n", resp.StatusCode)

	doc, err := goquery.NewDocument(url)
	if err != nil {
		panic(err)
	}

	var title, path string

	doc.Find("ol li").Each(func(i int, s *goquery.Selection) {
		if i == n {
			title = s.Find("a").Text()
			path, _ = s.Find("a").Attr("href")
		}
	})

	return title, url + path
}

func main() {
	// make random number
	rand.Seed(time.Now().UnixNano())
	var n int = rand.Intn(max)

	t, u := GetEssay(n)

	fmt.Printf("Essay No.%d: %s - %s\n", n, t, u)
}
