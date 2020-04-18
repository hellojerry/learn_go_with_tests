package main

import (
	"testing"
	"bytes"
	"reflect"
	"time"
)

const write = "write"
const sleep = "sleep"

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep(){
	s.Calls++
}

type SpyTime struct {
	durationSlept time.Duration
}

func(s *SpyTime) Sleep(duration time.Duration){
	s.durationSlept = duration
}

type CountdownOperationsSpy struct {
	Calls []string
}

func (s *CountdownOperationsSpy) Sleep(){
	s.Calls = append(s.Calls,sleep)
}

func(s *CountdownOperationsSpy) Write(
	p []byte) (n int, err error){
	s.Calls = append(s.Calls,write)
	return
}

func TestCountdown(t *testing.T){

	t.Run("sleep before every print",
		func(t *testing.T){
		spySleepPrinter := &CountdownOperationsSpy{}
		Countdown(spySleepPrinter, spySleepPrinter)
		res := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}
		if !reflect.DeepEqual(res,
			spySleepPrinter.Calls){
			t.Errorf("wanted calls %v got %v", res,spySleepPrinter.Calls)
		}

	})

	t.Run("prints 3 to Go!", func(t *testing.T){
		buffer := &bytes.Buffer{}
		Countdown(buffer, &CountdownOperationsSpy{})
		res := buffer.String()
		expected := `3
2
1
Go!`
	if res != expected {
		t.Errorf("res: %q, expected: %q",res, expected)
	}

	})
}

func TestConfigurableSleeper(t *testing.T){
	sleepTime := 5*time.Second
	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime,
			spyTime.Sleep}
	sleeper.Sleep()
	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime,spyTime.durationSlept)
	}

}
