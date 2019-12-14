package test

import (
	"io/ioutil"
	"strings"
	"testing"
	"unsafe"

	"github.com/yoshimitsuEgashira/97programmerbot/api"
)

func TestMakeText(t *testing.T) {
	text, err := api.MakeText()

	if err != nil {
		t.Errorf("Failed to make tweet text: %s", err.Error())
	}
	if text == "" {
		t.Error("Tweet text is empty.")
	}

	rows, err := ioutil.ReadFile("./data/essays.json")
	if err != nil {
		t.Errorf("Failed to open file %s", err.Error())
	}

	deleteWords := []string{"【 ", " 】", "著者："}
	for _, w := range deleteWords {
		text = strings.Replace(text, w, "", 1)
	}
	expect := *(*string)(unsafe.Pointer(&rows))
	actuals := strings.Split(text, "\n")
	for _, actual := range actuals {
		result := strings.Index(expect, actual)
		if result == -1 {
			t.Errorf("Could not match text in JSON data.\nexpect :%#v\nactual :%#v", expect, actual)
		}
	}
}
