package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

type Essay struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Link   string `json:"link"`
}

const hashTags = "#エッセイ #エンジニア #プログラマー #プログラマが知るべき97のこと"

func MakeText() (string, error) {
	var tweetText string
	rows, err := ioutil.ReadFile("./data/essays.json")
	if err != nil {
		return "", err
	}

	var essays []Essay
	err = json.Unmarshal(rows, &essays)
	if err != nil {
		fmt.Printf("Failed to parse %s\n", err.Error())
		return "", err
	}

	// generate random number
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(6)

	for i, e := range essays {
		if i == n {
			tweetText = "【 " + e.Title + " 】\n著者 : " + e.Author + "\n\n" + e.Link + "\n\n"
		}
	}
	fmt.Printf("Make Text Result : %s\n", tweetText+hashTags)
	return tweetText + hashTags, nil
}
