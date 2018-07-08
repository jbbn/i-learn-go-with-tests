package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const german = "German"
const helloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const germanHelloPrefix = "Halo, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := helloPrefix

	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case german:
		prefix = germanHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
