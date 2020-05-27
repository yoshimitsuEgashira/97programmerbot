package test

import (
	"fmt"
	"testing"

	"github.com/yossiee/97programmerbot/api"
)

func TestAuth(t *testing.T) {
	credentials := api.Auth()
	if credentials == nil {
		t.Errorf("Twitter verification failed.")
	}
	fmt.Println(credentials)
}
