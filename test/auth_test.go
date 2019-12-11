package test

import (
	"fmt"
	"testing"

	"github.com/yoshimitsuEgashira/97programmerbot/api"
)

func TestAuth(t *testing.T) {
	credentials := api.Auth()
	if credentials == nil {
		t.Errorf("Failed to twiiter authentication")
	}
	fmt.Println(credentials)
}
