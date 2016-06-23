package test

import (
	"testing"
)

func TestGreeting(t *testing.T) {
	expected := "hello"
	actual := greeting()
	if   expected != actual {
		t.Errorf("expected is [%s], actual is [%s]", expected, actual)
	}

}
