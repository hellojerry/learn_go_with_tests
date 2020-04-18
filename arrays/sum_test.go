package main

import (
	"testing"
	"reflect"
)
func TestSum(t *testing.T){

	assertCorrect := func(t *testing.T,res int, expected int, arr []int){
		t.Helper()
		if res != expected {
			t.Errorf("got %d, expected %d, given %v", res,expected,arr)
		}
	}

	t.Run("collection of 5 numbers",
		func(t *testing.T){
		numbers := []int{1,2,3,4,5}
		res := Sum(numbers)
		expected := 15
		assertCorrect(t,res,expected,numbers)
	})

	t.Run("collection of any size",
		func(t *testing.T){
		numbers := []int{1,2,3}
		res := Sum(numbers)
		expected := 6
		assertCorrect(t,res,expected,numbers)
	})
}

func TestSumAllTails(t *testing.T){

	checkSums := func(t *testing.T, res []int, expected []int){
		t.Helper()
		if !reflect.DeepEqual(res,expected){
			t.Errorf("res: %v, expected: %v",
				res, expected)
		}
	}

	t.Run("make the sums of some slices",
		func(t *testing.T){
		res := SumAllTails([]int{1,2}, []int{0,9})
		expected := []int{2,9}
		checkSums(t,res,expected)
	})
	t.Run("safely sum empty slices",
		func(t *testing.T){
		res := SumAllTails([]int{},
			[]int{3,4,5})
		expected := []int{0,9}
		checkSums(t,res,expected)
	})

}
