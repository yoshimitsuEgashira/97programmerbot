package test

import (
	"os"
	"testing"

	"github.com/yoshimitsuEgashira/97programmerbot/api"
)

func TestMakeText(t *testing.T) {
	err := os.Chdir("../")
	if err != nil {
		t.Errorf("Failed to change directory : %s", err.Error())
	}
	text, err := api.MakeText()

	if err != nil {
		t.Errorf("Failed to make tweet text: %s", err.Error())
	}
	if text == "" {
		t.Error("Tweet text is empty.")
	}
}
