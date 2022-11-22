package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	Countdown(buffer, spySleeper)

	result := buffer.String()
	expect := "3\n2\n1\nGo!"

	if result != expect {
		t.Errorf("\nresult: %q\nexpect: %q\n", result, expect)
	}

	if spySleeper.Calls != 3 {
		t.Errorf("\nwrong number of calls.\ncounted: %q\nexpected: 3\n", spySleeper.Calls)
	}
}
