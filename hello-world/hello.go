package main

import "fmt"

const	(
	spanish = "Spanish"
	french = "French"
	portuguese = "Portuguese"

	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix = "Bounjour, "
	portugueseHelloPrefix = "Olá, "
	englishHelloPrefix = "Hello, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case portuguese:
		prefix = portugueseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}