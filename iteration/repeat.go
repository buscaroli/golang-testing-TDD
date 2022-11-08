package iteration

// takes a string a an integer 'n' and returns the string repeated n-times.
func Repeat(s string, n int) string {
	var result string
	for i := 0; i < n; i++ {
		result += s
	}
	return result
}

// Takes a string, a separator string and an integer 'n' and returns the string repeated n-times.
func RepeatWithSeparator(str, sep string, n int) string {
	var result string
	for i := 0; i < n; i++ {
		result = result + str

		if i < n-1 {
			result += sep
		}
	}

	return result
}

// Takes a string, a separator string and an integer 'n' and returns the string repeated n-times. More performant version of RepeatWithSeparator()
func RepeatWithSeparator2(str, sep string, n int) string {
	var result string
	for i := 0; i < n; i++ {
		result = result + str

		if i < n-1 && sep != "" {
			result += sep
		}
	}

	return result
}
