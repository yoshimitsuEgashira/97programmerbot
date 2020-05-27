package main

import (
	"fmt"

	"github.com/yoshimitsuEgashira/97programmerbot/api"
)

func main() {
	// scrape

	// make tweet text
	t, err := api.MakeText()
	if err != nil {
		fmt.Printf("Failed to make tweet text %s\n", err.Error())
	}
	if t == "" {
		fmt.Println("Tweet text is empty")
		return
	}

	// twiter api auth
	apiAuth, err := api.Auth()
	if err != nil {
		return
	}

	// twitter post api
	_, err = api.PostTweet(*apiAuth, t, nil)
	if err != nil {
		fmt.Printf("Tweet failed : %s\n", err.Error())
		return
	}
	fmt.Println("tweeted successfully!")
}
