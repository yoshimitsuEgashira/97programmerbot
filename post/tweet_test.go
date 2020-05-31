package post

import (
	"testing"
	"time"

	. "github.com/yossiee/97programmerbot/auth"
)

func TestPostTweetSuccess(t *testing.T) {
	a := Auth()

	err := PostTweet(*a, time.Now().Format("2006年01月02日 15時04分05秒"), nil)
	if err != nil {
		t.Errorf("\nexpect: nil.\nactual: %#v.\n", err.Error())
	}
}

func TestPostTweetFail(t *testing.T) {
	a := Auth()

	err := PostTweet(*a, "", nil)
	if err == nil {
		t.Errorf("\nexpect: ERROR\nactual: %#v", err.Error())
	}
}
