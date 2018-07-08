package main

import "fmt"

const spanish = "Spanish"
const helloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == spanish {
		return spanishHelloPrefix + name
	}

	return helloPrefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
