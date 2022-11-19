package greeting

import (
	"fmt"
	"io"
)

// Sends a greeting message "Hello <name>" to the given io.Writer.
func Greet(w io.Writer, name string) {
	fmt.Fprintf(w, "Hello %s", name)
}
