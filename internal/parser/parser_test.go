package parser

import (
	"testing"
)

func TestNew(t *testing.T) {
	p := New()
	if p == nil {
		t.Fatal("expected parser, got nil")
	}
}

func TestLoadDirectory(t *testing.T) {
	p := New()
	// Add actual test once we have fixture data
	_ = p
}
