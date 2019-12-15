package main

import (
	"fmt"

	"github.com/yoshimitsuEgashira/97programmerbot/api"
)

func main() {
	apiAuth := api.Auth()

	t, err := api.MakeText()
	if err != nil {
		fmt.Printf("Failed to make tweet text %s\n", err.Error())
	}
	if t == "" {
		fmt.Println("Tweet text is empty")
		return
	}

	tw, err := api.PostTweet(*apiAuth, t, nil)
	if err != nil {
		fmt.Printf("Tweet failed : %s\n", err.Error())
		return
	}
	fmt.Printf("tweet is : %#v\n", tw)
}
