package greeting

import (
	"bytes"
	"testing"
)

// Dependency Injection consists of passing a dependency as a parameter (in this case io.Writer) so that we can:
// - test our code (otherwise how coud you test the a specific string was printed passing a string to the "Greet" function snd using fmt.Printf ?)
// - separate concerns: use dependency injection when it feels like your function is doing too many things (eg creating data and sending data (in this case we can choose where to send the data))
// - reuse our code

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	// bytes.Buffer implements io.Writer
	// io.Writer is implemented by many packages such as:
	// - http.ResponseWriter
	// - fmt.Sprintf
	Greet(&buffer, "Matt")

	result := buffer.String()
	expect := "Hello Matt"

	if result != expect {
		t.Errorf("\nresult: %s\nexpect: %s\n", result, expect)
	}
}
