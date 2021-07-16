package main

import "fmt"

const spanish = "Spanish"
const world = "World"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

func Hello(name string, language string) string {
	if name == "" {
		name = world
	}

	if language == spanish {
		return spanishHelloPrefix + name
	} else {
		return englishHelloPrefix + name
	}
}

func main() {
	fmt.Println(Hello("world", ""))
}