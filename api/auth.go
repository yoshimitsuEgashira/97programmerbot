package api

import (
	"fmt"
	"log"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/joho/godotenv"
)

func Auth() (*anaconda.TwitterApi, error) {
	err := godotenv.Load(fmt.Sprintf("./%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		log.Printf("ERROR: %#v", err.Error())
		return nil, err
	}

	anaconda.SetConsumerKey(os.Getenv("CONSUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("CONSUMER_SECRET"))
	return anaconda.NewTwitterApi(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET")), err
}
