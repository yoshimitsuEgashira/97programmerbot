package api

import (
	"fmt"

	"github.com/ChimeraCoder/anaconda"
)

func PostTweet(api anaconda.TwitterApi, text string, urls map[string][]string) (anaconda.Tweet, error) {
	tw, err := api.PostTweet(text, urls)
	if err != nil {
		fmt.Printf("Failed to tweet : %s\n", err.Error())
		return tw, err
	}
	return tw, nil
}
