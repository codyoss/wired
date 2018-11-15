package http_test

import (
	"fmt"
	"testing"

	"github.com/codyoss/wired/pkg/rpc/http"
)

func TestErrorMessageIsError(t *testing.T) {
	// This test asserts that ErrorMessage fullfills the error interface. This would not compile if not true.
	var _ error = (*http.ErrorMessage)(nil)
}

func TestErrorMessage(t *testing.T) {
	str := "bad things"
	em := http.NewErrorMessage(fmt.Errorf(str))
	got := em.Error()
	if got != str {
		t.Errorf("got %v, want %v", got, str)
	}
}
