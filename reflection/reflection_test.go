package main

import (
	"reflect"
	"testing"
)


type Profile struct {
	Age int
	City string

}

type Person struct {
	Name string
	Profile Profile
}


func assertContains(t *testing.T, haystack []string,
		needle string){

	contains := false
	for _,x := range haystack{
		if x == needle {

			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q",haystack,needle)
	}
}

func TestWalk(t *testing.T){

	cases := []struct{
		Name string
		Input interface{}
		ExpectedCalls []string
	} {
		{
			"Struct with one string field",
			struct {
				Name string
			} {"Chris"},
			[]string{"Chris"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris","London"},
			[]string{"Chris","London"},
		},
		{
			"Struct with non-string field",
			struct{
				Name string
				Age int
			}{"Chris",33},
			[]string{"Chris"},
		},
		{
			"Nested Fields",
			Person{
				"Chris",
				Profile{33,"London"},
			},
			[]string{"Chris","London"},
		},
		{
			"Pointers to things",
			&Person{
				"Chris",
				Profile{33,"London"},
			},
			[]string{"Chris","London"},
		},
		{
			"Slices",
			[]Profile {
				{33,"London"},
				{34,"Reykjavik"},
			},
			[]string{"London","Reykjavik"},
		},
		{
			"Arrays",
			[2]Profile {
				{33,"London"},
				{34,"Reykjavik"},
			},
			[]string{"London","Reykjavik"},
		},
	}
	for _,test := range cases {
		t.Run(test.Name, func(t *testing.T){
			var res []string
			walk(test.Input, func(input string){				res = append(res,input)
			})
			if !reflect.DeepEqual(res,test.ExpectedCalls){
				t.Errorf("got %v, want %v", res, test.ExpectedCalls)
			}

		})
	}

	t.Run("with maps", func(t *testing.T){
		aMap := map[string]string{
			"Foo":"Bar",
			"Baz":"Boz",
		}
		var res []string
		walk(aMap,func(input string){
			res = append(res,input)
		})
		assertContains(t,res,"Bar")
		assertContains(t,res,"Boz")
	})
	t.Run("with channels", func(t *testing.T){
		aChannel := make(chan Profile)
		go func(){
			aChannel <- Profile{33,"Berlin"}
			aChannel <- Profile{34,"Katowice"}
			close(aChannel)
		}()
		var res []string
		expected := []string{"Berlin","Katowice"}
		walk(aChannel, func(input string){
			res = append(res,input)
		})
		if !reflect.DeepEqual(res,expected){
			t.Errorf("got %v, want %v",res,expected)

		}
	})
	t.Run("With function",func(t *testing.T){
		aFunction := func()(Profile,Profile){
			return Profile{33,"Berlin"}, Profile{34,"Katowice"}
		}
		var res []string
		expected := []string{"Berlin","Katowice"}
		walk(aFunction, func(input string){
			res = append(res,input)
		})
		if !reflect.DeepEqual(res,expected){
			t.Errorf("got %v, want %v",res,expected)
		}
	})

}
