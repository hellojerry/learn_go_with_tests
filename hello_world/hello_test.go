package main

import "testing"

func TestHello(t *testing.T){

	assertCorrectMessage := func(t *testing.T, res string, expected string){
		t.Helper()
		if res != expected {
			t.Errorf("received %q, expected %q", res, expected)
		}
	}

	t.Run("saying hello to people",
		func(t *testing.T){

		res := Hello("mike","")
		expected := "hello mike"
		assertCorrectMessage(t,res,expected)
	})

	t.Run("Say 'hello world' when an empty string is supplied", func(t *testing.T){

		res := Hello("","")
		expected := "hello world"
		assertCorrectMessage(t,res,expected)
	})

	t.Run("saying hello in spanish",
		func(t *testing.T){
		res := Hello("mike","Spanish")
		expected := "hola mike"
		assertCorrectMessage(t,res,expected)

	})
	t.Run("saying hello in french",
		func(t *testing.T){
			res := Hello("mike","French")
			expected := "bonjour mike"
			assertCorrectMessage(t,res,expected)
	})
}
