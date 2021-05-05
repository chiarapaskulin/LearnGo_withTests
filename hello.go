package main

import "fmt"

// Lesson 1
// func main() {
// 	fmt.Println("Hello, World")
// }

// Lesson 2
// func Hello() string {
// 	return "Hello, world"
// }

// func main() {
// 	fmt.Println(Hello("world"))
// }

// Lesson 3
// func Hello(name string) string {
// 	return "Hello, " + name
// }

// const englishHelloPrefix = "Hello, "

// func Hello(name string) string {
// 	return englishHelloPrefix + name
// }

func main() {
	fmt.Println(Hello("world", ""))
}

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const portugueseHelloPrefix = "Ola, "
const spanish = "Spanish"
const french = "French"
const portuguese = "Portuguese"

func Hello(name, language string) string {
	if name == "" {
		name = "world"
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

// Lesson 5
// func Hello(name, language string) string {
// 	if name == "" {
// 		name = "world"
// 	}
// 	if language == spanish {
// 		return spanishHelloPrefix + name
// 	}
// 	if language == french {
// 		return frenchHelloPrefix + name
// 	}

// 	return englishHelloPrefix + name
// }
