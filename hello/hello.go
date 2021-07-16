package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const world = "World"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

// Hello public function: starts with uppercase letter
func Hello(name string, language string) string {
	if name == "" {
		name = world
	}
	return greetingPrefix(language) + name
}

// private function: starts with lowercase letter
func greetingPrefix(language string) (prefix string) {	// prefix defaults to ""
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}