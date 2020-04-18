package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T){
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")
	res := buffer.String()
	expected := "Hello, Chris"

	if res != expected {
		t.Errorf("res %q, expected %q",
			res,expected)
	}

}
