package api

import (
	"fmt"

	"github.com/ChimeraCoder/anaconda"
)

func PostTweet(api anaconda.TwitterApi, text string, urls map[string][]string) (anaconda.Tweet, error) {
	tw, err := api.PostTweet(text, urls)
	if err != nil {
		fmt.Println(err.Error())
		return tw, err
	}
	fmt.Println(tw)
	return tw, nil
}
