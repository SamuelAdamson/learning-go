package hello

import (
	"testing"
)

func TestHello_WithArray(t *testing.T) {
	var want string = "Hello - test1, test2, test3!"
	var got string = Hello([]string{"test1", "test2", "test3"})

	if want != got {
		t.Errorf("Wanted '%s' got '%s'", want, got)
	}
}

func TestHello_WithEmptyArray(t *testing.T) {
	var want string = "Hello... no one :("
	var got string = Hello([]string{})

	if want != got {
		t.Errorf("Wanted '%s' got '%s'", want, got)
	}
}
