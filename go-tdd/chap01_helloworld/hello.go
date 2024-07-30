package main

import "fmt"

const englishHelloPrefix = "Hello, "
const frenchHelloPrefix = "Bonjour, "
const spanishHelloPrefix = "Hola, "
const spanish = "Spanish"
const french = "French"

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(name, language)
}

func greetingPrefix(name string, language string) string {
	switch language {
	case french:
		return frenchHelloPrefix + name
	case spanish:
		return spanishHelloPrefix + name
	default:
		return englishHelloPrefix + name
	}
}

func main() {
	fmt.Println(Hello("World", ""))
}
