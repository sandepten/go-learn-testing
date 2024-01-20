package main

const (
	spanish = "Spanish"
	french  = "French"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

// public function as it starts with uppercase
func Hello(name, lang string) string {
	if name == "" {
		name = "World"
	}
	return getPrefix(lang) + name
}

// by giving a name to the return value, a naked return is possible which will return the value of prefix
// private function as it starts with lowercase
func getPrefix(lang string) (prefix string) {
	switch lang {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
