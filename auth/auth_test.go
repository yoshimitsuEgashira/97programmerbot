package auth

import (
	"testing"
)

const env = "../%s.env"

func TestAuthSuccess(t *testing.T) {
	c, err := Auth(env)
	if err != nil {
		t.Errorf("\nexpect: nil.\nactual: %#v", err.Error())
	}
	if c == nil {
		t.Errorf("\nexpect: %#v.\nactual: nil", c)
	}
}

func TestAuthFail(t *testing.T) {
	_, err := Auth("")
	if err == nil {
		t.Errorf("\nexpect: ERROR.\nactual: %#v", err)
	}
}
