package main

import "testing"

func TestCountWords(t *testing.T) {
	got := countWords("Go make small CLI tools easy")

	want := 6

	if got != want {
		t.Errorf("countWords() = %d, want %d", got, want)
	}
}
