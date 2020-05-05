package api

import "github.com/ChimeraCoder/anaconda"

func PostTweet(api anaconda.TwitterApi, text string, urls map[string][]string) (anaconda.Tweet, error) {
	tw, err := api.PostTweet(text, urls)
	if err != nil {
		return tw, err
	}
	return tw, nil
}
