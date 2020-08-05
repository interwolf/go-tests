package main

import "fmt"

func main() {
	fmt.Println(Hello("Yiming", "cn") + ", welcom to go world :-)")
}

// HelloPrefix prepares prefix for hello
func HelloPrefix(language string) (helloLang string) {
	switch language {
	case "es":
		helloLang = "Hola"
	case "fr":
		helloLang = "Bonjour"
	case "cn":
		helloLang = "Nihao"
	case "jp":
		helloLang = "Soga"
	case "en":
		helloLang = "Hello"
	default:
		helloLang = "no_lang"
	}

	return
}

// Hello prints hello in different languages
func Hello(name string, language string) string {
	prefix := HelloPrefix(language)
	if prefix == "no_lang" {
		return "Haha, world"
	}
	return prefix + ", " + name
}
