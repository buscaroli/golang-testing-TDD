package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestCountdown(t *testing.T) {
	t.Run("tests Sleep is called three times", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpySleeper{}

		Countdown(buffer, spySleeper)

		result := buffer.String()
		expect := "3\n2\n1\nGo!"

		if result != expect {
			t.Errorf("\nresult: %q\nexpect: %q\n", result, expect)
		}
	})

	t.Run("tests the order of sleeps and prints: first sleep, then print", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperations{}
		// Replace Countdown with LogicallyWrongCountdown to see the test failing
		Countdown(spySleepPrinter, spySleepPrinter)

		expect := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(expect, spySleepPrinter.Calls) {
			t.Errorf("\nexpected: %v\ngot: %v\n", expect, spySleepPrinter.Calls)
		}
	})
}
