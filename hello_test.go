package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, World!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHelloUser(t *testing.T) {
	t.Run("saying 'Hello Username', when a name is provided", func(t *testing.T) {
		got := HelloUser("Matt")
		want := "Hello, Matt!"

		assertCorrectMsg(t, got, want)

	})

	t.Run("Saying 'Hello there' when the username is not provided", func(t *testing.T) {
		got := HelloUser("")
		want := "Hello there!"

		assertCorrectMsg(t, got, want)
	})

}

// For helper functions, it's a good idea to accept a testing.TB which is an interface that *testing.T and *testing.B both satisfy, so you can call helper functions from a test, or a benchmark
// t.Helper() is needed to tell the test suite that this method is a helper. By doing this when it fails the line number reported will be in our function call rather than inside our test helper.
func assertCorrectMsg(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("\ngot: %q\nwant: %q", got, want)
	}
}
