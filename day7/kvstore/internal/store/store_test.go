package store

import "testing"

func TestSetAndGet(t *testing.T) {
	Set("username", "alice")
	got, ok := Get("username")

	if !ok {
		t.Fatal("expected key to exist")
	}
	if got != "alice" {
		t.Errorf("expected 'alice', got '%s'", got)
	}
}
