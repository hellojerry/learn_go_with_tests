package integers

import (
	"testing"
	"fmt"
)


func TestAdder(t *testing.T){

	assertCorrectResult := func(t *testing.T,
			res int, expected int){
		t.Helper()
		if res != expected {
			t.Errorf("expected '%d' but received '%d'",expected,res)
		}
	}
	t.Run("adding two numbers", func(t *testing.T){
		res := Add(2,2)
		expected := 4
		assertCorrectResult(t,res,expected)
	})
}
func ExampleAdd(){
	sum := Add(1,5)
	fmt.Println(sum)
	// Output: 6
}
