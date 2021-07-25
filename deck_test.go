package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("error1")

	}
	if d[0] != "ace" {
		t.Errorf("error2")
	}
}
