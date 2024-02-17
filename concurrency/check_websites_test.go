package main

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	return url != "waat://foobar.baz"
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"https://google.com",
		"https://eckertalex.dev",
		"waat://foobar.baz",
	}

	want := map[string]bool{
		"https://google.com":     true,
		"https://eckertalex.dev": true,
		"waat://foobar.baz":      false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("wanted %v got %v", want, got)
	}
}
