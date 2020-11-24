package main

import "testing"

func TestIndex(t *testing.T) {
	result := setText("Teste: Code.education Rocks!")
	response := "<b>Teste</b>"

	if result != response {
		t.Errorf("Invalid result! :( return %s, wanted %s", result, response)
	}
}
