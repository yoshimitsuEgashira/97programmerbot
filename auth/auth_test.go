package auth

import (
	"testing"
)

func TestAuth(t *testing.T) {
	credentials := Auth()

	if credentials == nil {
		t.Errorf("\nexpect: nil.\nactual: %#v", credentials)
	}
}
