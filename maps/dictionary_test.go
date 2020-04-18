package main

import (
	"testing"
)

func assertStrings(t *testing.T,res,expected string){

	t.Helper()
	if res != expected {
		t.Errorf("res: %q, expected: %q", res,expected)

	}
}

func assertError(t *testing.T, res, expected error){

	t.Helper()
	if res != expected{
		t.Errorf("got error %q wanted error %q", res,expected)
	}
	if res == nil {
		if expected == nil {
			return
		}
		t.Fatal("expected to get an error")
	}

}

func assertDefinition(t *testing.T, dictionary Dictionary,word, expected string){
	t.Helper()
	res,err := dictionary.Search(word)
	if err != nil{
		t.Fatal("should find added word:",err)
	}
	assertStrings(t,res,expected)
}

func TestSearch(t *testing.T){
	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("known word", func(t *testing.T){
		res,_ := dictionary.Search("test")
		expected := "this is just a test"
		assertStrings(t,res,expected)
	})

	t.Run("unknown word", func(t *testing.T){
		_, err := dictionary.Search("unknown")
		assertError(t,err,ErrNotFound)
	})

}

func TestAdd(t *testing.T){
	t.Run("new word", func(t *testing.T){
	dictionary := Dictionary{}
	err := dictionary.Add("test", "this is just a test")
	expected := "this is just a test"
	assertError(t,err,nil)
	assertDefinition(t,dictionary,"test",expected)
	})
	t.Run("existing word", func(t *testing.T){
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word:definition}
		err := dictionary.Add(word,"new test")
		assertError(t,err,ErrWordExists)
		assertDefinition(t,dictionary,word,definition)
	})
}

func TestUpdate(t *testing.T){
	t.Run("existing word", func(t *testing.T){
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word:definition}
		newDefinition := "new definition"
		err := dictionary.Update(word,newDefinition)
		assertError(t,err,nil)
		assertDefinition(t,dictionary,word,newDefinition)
	})

	t.Run("new word", func(t *testing.T){
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}
		err := dictionary.Update(word,definition)
		assertError(t,err,ErrWordDoesNotExist)

	})

}

func TestDelete(t *testing.T){

	word := "test"
	dictionary := Dictionary{word:"test definition"}
	dictionary.Delete(word)
	_, err := dictionary.Search(word)
	if err != ErrNotFound {
		t.Errorf("expected %q to be deleted", word)
	}

}
