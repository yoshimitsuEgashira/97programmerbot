package auth

import (
	"os"

	a "github.com/ChimeraCoder/anaconda"
)

func Auth() *a.TwitterApi {
	a.SetConsumerKey(os.Getenv("CONSUMER_KEY"))
	a.SetConsumerSecret(os.Getenv("CONSUMER_SECRET"))
	return a.NewTwitterApi(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET"))
}
