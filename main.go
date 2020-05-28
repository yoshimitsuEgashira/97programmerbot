package main

import (
	"log"
	"math/rand"
	"time"

	. "github.com/yossiee/97programmerbot/auth"
	. "github.com/yossiee/97programmerbot/post"
	. "github.com/yossiee/97programmerbot/scraper"
)

const (
	env = "./%s.env"
	url = "https://xn--97-273ae6a4irb6e2hsoiozc2g4b8082p.com"
	max = 106
)

func main() {
	rand.Seed(time.Now().UnixNano())
	t, u, err := FetchEssay(rand.Intn(max), url)
	if err != nil {
		return
	}

	a, err := Auth(env)
	if err != nil {
		log.Printf("ERROR: %#v\n", err.Error())
		return
	}

	text := "【 " + t + " 】\n\n" + u
	err = PostTweet(*a, text, nil)
	if err != nil {
		log.Printf("Failed to tweet : %s\n", err.Error())
		return
	}
	log.Println("tweeted successfully!")
}
