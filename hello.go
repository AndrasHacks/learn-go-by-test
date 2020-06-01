package main

import "fmt"

func Hello(name string) string {
	if len(name) == 0 {
		return "Hello Gophers!"
	}
	return "Hello " + name + "!"
}

func main() {
	fmt.Printf(Hello(""))
}
