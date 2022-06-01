package utils

import (
	"reflect"
	"testing"
)

func TestPopulate(t *testing.T) {
	wordlist := []string{"test1", "test2"}
	got := Populate(wordlist)
	want := map[string]string{
		"test1": "111110",
		"test2": "111111",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
