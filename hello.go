package main

import (
	"fmt"
)

// Constants should improve performance of your application as it saves you creating the "Hello, " string instance every time Hello is called.
const HelloPrefix = "Hello "
const HolaPrefix = "Hola "
const SalutPrefix = "Salut "
const spanish = "spanish"
const french = "french"

func Hello() string {
	return "Hello, World!"
}

func HelloUser(username string, language string) string {
	if username == "" {
		username = ""
	}

	return greetingPrefix(language) + username + "!"
}

func greetingPrefix(language string) string {
	var greeting string

	switch language {
	case spanish:
		greeting = HolaPrefix
	case french:
		greeting = SalutPrefix
	default:
		greeting = HelloPrefix
	}

	return greeting
}

func main() {
	fmt.Println(Hello())
	fmt.Println(HelloUser("Matt", ""))
}
