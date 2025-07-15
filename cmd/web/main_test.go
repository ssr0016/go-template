package main

import (
	"testing"
)

func TestRunt(t *testing.T) {
	err := run()
	if err != nil {
		t.Error("failed to run app", err)
	}
}
