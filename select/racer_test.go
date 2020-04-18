package main

import (
	"testing"
	"time"
	"net/http"
	"net/http/httptest"
)

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))

}

func TestRacer(t *testing.T){
	t.Run("compares speeds of two servers, returning url of the fastest one", func(t *testing.T){
	slowServer := makeDelayedServer(20*time.Millisecond)
	fastServer := makeDelayedServer(0*time.Millisecond)
	defer slowServer.Close()
	defer fastServer.Close()
	slowURL := slowServer.URL
	fastURL := fastServer.URL

	expected := fastURL
	res,err := Racer(slowURL,fastURL)

	if err != nil{
		t.Fatalf("did not expect an error but got one %v",err)
	}

	if res != expected {
		t.Errorf("got %q, wanted %q",
			res,expected)
	}


	})

	t.Run("returns an error if a server doesnt respond within 10s", func(t *testing.T){
		serverA := makeDelayedServer(25*time.Millisecond)
		defer serverA.Close()
		_,err := Racer(serverA.URL,serverA.URL)
		if err != nil{
			t.Error("expected an error but didnt get one")
		}

	})

}
