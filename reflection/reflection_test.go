package main

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"Struct with non string fields",
			struct {
				Name string
				City string
				Age  int
			}{"Chris", "London", 33},
			[]string{"Chris", "London"},
		},
		{
			"Nested struct",
			Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"Pointer to things",
			&Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}
