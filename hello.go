package main

import "fmt"

func Hello(name string, language string) string {
	const spanish = "Spanish"
	const french = "French"
	const englishHelloPrefix = "Hello, "
	const spanishHelloPrefix = "Hola, "
	const frenchHelloPrefix = "Bonjur, "
	if len(name) == 0 {
		name = "World"
	}
	if language == french {
		return frenchHelloPrefix + name + "!"
	}

	if language == spanish {
		return spanishHelloPrefix + name + "!"
	}
	return englishHelloPrefix + name + "!"
}

func main() {
	fmt.Printf(Hello("", ""))
}
