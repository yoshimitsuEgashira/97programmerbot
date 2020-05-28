package post

import (
	"testing"

	. "github.com/yossiee/97programmerbot/auth"
)

const env = "../%s.env"

func TestPostTweetSuccess(t *testing.T) {
	a, _ := Auth(env)

	err := PostTweet(*a, "Test tweet", nil)
	if err != nil {
		t.Errorf("\nexpect: nil.\nactual: %#v.\n", err.Error())
	}
}

func TestPostTweetFail(t *testing.T) {
	a, _ := Auth(env)

	err := PostTweet(*a, "", nil)
	if err == nil {
		t.Errorf("\nexpect: ERROR\nactual: %#v", err.Error())
	}
}
