package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "hello "
const spanishHelloPrefix = "hola "
const frenchHelloPrefix = "bonjour "

func Hello(name string, language string) string {
	if len(name) < 1 {
		name = "world"
	}
	prefix := ""
	switch language {
		case spanish:
			prefix = spanishHelloPrefix
		case french:
			prefix = frenchHelloPrefix
		default:
			prefix = englishHelloPrefix
	}
	return prefix + name
}


func main(){
	fmt.Println(Hello("mike",""))
}
