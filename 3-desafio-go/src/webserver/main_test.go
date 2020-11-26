package main

import "testing"

func TestIndex(t *testing.T) {
	result := setText("Code.education Rocks!")
	response := "<b>Code.education Rocks!</b>"

	if result != response {
		t.Errorf("Invalid result! :( return %s, wanted %s", result, response)
	}
}