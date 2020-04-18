package concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool{

	if url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T){

	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}
	expected := map[string]bool{

		"http://google.com": true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds": false,
	}
	res := CheckWebsites(mockWebsiteChecker,
			websites)
	if !reflect.DeepEqual(res,expected){
		t.Fatalf("wanted %v, got %v",
			res,expected)
	}
}