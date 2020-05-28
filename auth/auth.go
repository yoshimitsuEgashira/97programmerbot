package auth

import (
	"fmt"
	"os"

	a "github.com/ChimeraCoder/anaconda"
	g "github.com/joho/godotenv"
)

func Auth(path string) (*a.TwitterApi, error) {
	err := g.Load(fmt.Sprintf(path, os.Getenv("GO_ENV")))
	if err != nil {
		return nil, err
	}

	a.SetConsumerKey(os.Getenv("CONSUMER_KEY"))
	a.SetConsumerSecret(os.Getenv("CONSUMER_SECRET"))
	return a.NewTwitterApi(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET")), err
}
