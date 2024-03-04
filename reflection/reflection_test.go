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
	City string
	Age  int
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Bilbo"},
			[]string{"Bilbo"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Bilbo", "Hobbiton"},
			[]string{"Bilbo", "Hobbiton"},
		},
		{
			"struct with non string fields",
			struct {
				Name string
				Age  int
			}{"Bilbo", 111},
			[]string{"Bilbo"},
		},
		{
			"nested fields",
			Person{
				"Bilbo",
				Profile{"Hobbiton", 111},
			},
			[]string{"Bilbo", "Hobbiton"},
		},
		{
			"pointer to things",
			&Person{
				"Bilbo",
				Profile{"Hobbiton", 111},
			},
			[]string{"Bilbo", "Hobbiton"},
		},
		{
			"slices",
			[]Profile{
				{"Hobbiton", 111},
				{"Minas Tirith", 89},
			},
			[]string{"Hobbiton", "Minas Tirith"},
		},
		{
			"arrays",
			[2]Profile{
				{"Hobbiton", 111},
				{"Minas Tirith", 89},
			},
			[]string{"Hobbiton", "Minas Tirith"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string

			Walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("maps", func(t *testing.T) {
		aMap := map[string]string{
			"Cat": "Meow",
			"Dog": "Woof",
		}

		var got []string
		Walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Meow")
		assertContains(t, got, "Woof")
	})

	t.Run("channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{"Hobbiton", 111}
			aChannel <- Profile{"Minas Tirith", 89}
			close(aChannel)
		}()

		var got []string
		want := []string{"Hobbiton", "Minas Tirith"}

		Walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("functions", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{"Hobbiton", 111}, Profile{"Minas Tirith", 89}
		}

		var got []string
		want := []string{"Hobbiton", "Minas Tirith"}

		Walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}

	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
	}
}
