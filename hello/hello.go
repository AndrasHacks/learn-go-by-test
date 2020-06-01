package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjur, "

func Hello(name string, language string) string {
	if len(name) == 0 {
		name = "World"
	}

	return greetingPrefix(language) + name + "!"
}

// Introduction of a named return value!
// Private functions starts with lowercase letters.
func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return prefix
}

func main() {
	fmt.Printf(Hello("", ""))
}
