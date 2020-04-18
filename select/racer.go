package main

import (
	"time"
	"net/http"
	"fmt"
)

var tenSecondTimeout = 10*time.Second

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)

}

func ping(url string) chan struct{}{
	ch := make(chan struct{})
	go func(){
		http.Get(url)
		close(ch)
	}()
	return ch
}

func ConfigurableRacer(a,b string, timeout time.Duration) (winner string, err error){

	// whichever channel returns FIRST
	// is the winner. both pings
	// are started at the same time.
	select {
		case <-ping(a):
			return a,nil
		case <-ping(b):
			return b,nil
		case <-time.After(timeout):
			return "",fmt.Errorf("timed out waiting for %s and %s",a,b)
	}

}

func Racer(a,b string) (winner string, err error){
	return ConfigurableRacer(a,b,tenSecondTimeout)


}

