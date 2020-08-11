package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {

	testCases := []struct {
		Name     string
		Input    interface{}
		Expected []string
	}{
		{
			"One string",
			struct {
				Name string
			}{"Yiming"},
			[]string{"Yiming"},
		},
		{
			"Two strings",
			struct {
				Name string
				City string
			}{"Yiming", "Changsha"},
			[]string{"Yiming", "Changsha"},
		},
		{
			"Two strings and one int",
			struct {
				Name   string
				City   string
				Number int
			}{"Yiming", "Changsha", 1},
			[]string{"Yiming", "Changsha"},
		},
		{
			"Nested field",
			Person{
				"Yiming",
				Profile{33, "Changsha"},
			},
			[]string{"Yiming", "Changsha"},
		},
		{
			"Pointer to struct",
			&Person{
				"Yiming",
				Profile{33, "Changsha"},
			},
			[]string{"Yiming", "Changsha"},
		},
		{
			"Slice",
			[]Profile{
				{33, "Changsha"},
				{34, "Jinan"},
			},
			[]string{"Changsha", "Jinan"},
		},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			Walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.Expected) {
				t.Errorf("got: %v, expected: %v\n", got, test.Expected)
			}
		})
	}

}
