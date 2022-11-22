package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}

	Countdown(buffer)

	result := buffer.String()
	expect := "3\n2\n1\nGo!"

	if result != expect {
		t.Errorf("\nresult: %q\nexpect: %q\n", result, expect)
	}
}
