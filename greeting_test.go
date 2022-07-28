package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)

	msg, err := Hello(name)
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys")=%q,%v,want match for %#q,nill`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	hello, err := Hello("")
	if hello != "" || err == nil {
		t.Fatalf(`Hello("")=%q,%v,want "",error`, hello, err)

	}
}
