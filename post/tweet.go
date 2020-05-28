package post

import a "github.com/ChimeraCoder/anaconda"

func PostTweet(api a.TwitterApi, text string, urls map[string][]string) error {
	_, err := api.PostTweet(text, urls)
	if err != nil {
		return err
	}
	return nil
}
