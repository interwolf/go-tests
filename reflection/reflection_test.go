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
		{
			"Array",
			[2]Profile{
				{33, "Changsha"},
				{34, "Jinan"},
			},
			[]string{"Changsha", "Jinan"},
		},
		// {
		// 	"Map",
		// 	map[string]string{
		// 		"Yiming1": "B",
		// 		"Yiming2": "A",
		// 	},
		// 	[]string{"A", "B"},
		// },
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

	t.Run("Map no order", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "B",
			"Bar": "A",
		}

		var got []string
		Walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContain(t, got, "B")
		assertContain(t, got, "A")
	})

	t.Run("Channel", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "Changsha"}
			aChannel <- Profile{43, "Jinan"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Changsha", "Jinan"}

		Walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Func", func(t *testing.T) {
		aFunc := func() (Profile, Profile) {
			return Profile{33, "Changsha"}, Profile{43, "Jinan"}
		}

		var got []string
		want := []string{"Changsha", "Jinan"}

		Walk(aFunc, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContain(t *testing.T, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
			break
		}
	}

	if !contains {
		t.Errorf("expected %+v to contain %q but not", haystack, needle)
	}
}
