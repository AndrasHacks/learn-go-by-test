package main

import "fmt"

func Hello(name string) string {
	const englishHelloPrefix = "Hello, "
	if len(name) == 0 {
		return englishHelloPrefix + "World!"
	}
	return englishHelloPrefix + name + "!"
}

func main() {
	fmt.Printf(Hello(""))
}
