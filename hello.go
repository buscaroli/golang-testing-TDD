package main

import "fmt"

// Constants should improve performance of your application as it saves you creating the "Hello, " string instance every time Hello is called.
const HelloPrefix = "Hello, "
const HolaPrefix = "Hola, "
const spanish = "spanish"

func Hello() string {
	return "Hello, World!"
}

func HelloUser(u string, l string) string {
	if u == "" {
		if l == spanish {
			return "Hola!"
		}
		return "Hello there!"

	} else {
		if l == spanish {
			return HolaPrefix + u + "!"
		}
	}

	return HelloPrefix + u + "!"
}

func main() {
	fmt.Println(Hello())
	fmt.Println(HelloUser("Matt", ""))
}
