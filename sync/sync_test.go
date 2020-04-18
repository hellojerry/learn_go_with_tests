package main

import (
	"testing"
	"sync"
)

func assertCounter(t *testing.T,res *Counter, expected int){
	t.Helper()
	if res.Value() != expected{
		t.Errorf("got %d, want %d", res.Value(),3)
	}

}

func TestCounter(t *testing.T){
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T){
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()
		assertCounter(t,&counter,3)
	})
	t.Run("it runs safely concurrently", func(t *testing.T){
		wantedCount := 1000
		counter := Counter{}
		var wg sync.WaitGroup
		wg.Add(wantedCount)
		for i := 0; i < wantedCount; i++{
			go func(w *sync.WaitGroup){
				counter.Inc()
				w.Done()
			}(&wg)
		}
		wg.Wait()
		assertCounter(t,&counter,wantedCount)
	})
}
