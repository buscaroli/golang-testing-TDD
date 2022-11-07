package main

import "fmt"

// Constants should improve performance of your application as it saves you creating the "Hello, " string instance every time Hello is called.
const HelloPrefix = "Hello, "

func Hello() string {
	return "Hello, World!"
}

func HelloUser(u string) string {
	if u == "" {
		return "Hello there!"
	}
	return HelloPrefix + u + "!"
}

func main() {
	fmt.Println(Hello())
	fmt.Println(HelloUser("Matt"))
}
