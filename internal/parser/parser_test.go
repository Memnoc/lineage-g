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

/*
* FIX: fixture data placeholder
 */
func TestLoadDirectory(t *testing.T) {
	p := New()
	// this should contain fixture data
	_ = p
}
