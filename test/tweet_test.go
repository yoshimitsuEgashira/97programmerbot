package test

import (
	"testing"
	"time"

	"github.com/yoshimitsuEgashira/97programmerbot/api"
)

func TestPostTweet(t *testing.T) {
	apiAuth := api.Auth()
	tweetTime := time.Now()
	timeStr := tweetTime.Format("2006年01月02日 15時04分05秒")
	_, err := api.PostTweet(*apiAuth, timeStr+"\ntest text", nil)
	if err != nil {
		t.Errorf("Failed to post tweet : %s", err.Error())
	}
}
