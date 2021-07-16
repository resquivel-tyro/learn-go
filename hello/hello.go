package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const world = "World"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = world
	}

	if language == spanish {
		return spanishHelloPrefix + name
	} else if language == french {
		return frenchHelloPrefix + name
	} else {
		return englishHelloPrefix + name
	}
}

func main() {
	fmt.Println(Hello("world", ""))
}