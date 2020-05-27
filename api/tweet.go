package api

import (
	"log"

	"github.com/ChimeraCoder/anaconda"
)

func PostTweet(api anaconda.TwitterApi, text string, urls map[string][]string) error {
	_, err := api.PostTweet(text, urls)
	if err != nil {
		log.Printf("Failed to tweet : %s\n", err.Error())
		return err
	}
	return err
}
