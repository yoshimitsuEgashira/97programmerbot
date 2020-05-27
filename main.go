package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/yossiee/97programmerbot/api"
	"github.com/yossiee/97programmerbot/scraper"
)

const (
	url string = "https://xn--97-273ae6a4irb6e2hsoiozc2g4b8082p.com"
	max int    = 106
)

func main() {
	rand.Seed(time.Now().UnixNano())
	t, u, err := scraper.FetchEssay(rand.Intn(max), url)
	if err != nil {
		return
	}

	apiAuth, err := api.Auth()
	if err != nil {
		return
	}

	text := "【 " + t + " 】\n\n" + u
	err = api.PostTweet(*apiAuth, text, nil)
	if err != nil {
		return
	}
	fmt.Println("tweeted successfully!")
}
