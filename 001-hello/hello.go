package hello

import "fmt"

const enPrefix = "Hello, "
const esPrefix = "Hola, "
const frPrefix = "Bonjour, "

func greetPrefix(lang string) (prefix string) {
	switch lang {
	case "es":
		prefix = esPrefix
	case "fr":
		prefix = frPrefix
	default:
		prefix = enPrefix
	}
	return
}

func Hello(input, lang string) string {
	prefix := greetPrefix(lang)

	name := input
	if input == "" {
		name = "World"
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("", "en"))
}
