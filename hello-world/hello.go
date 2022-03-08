package hello

import "fmt"

const spanish = "espanhol"
const french = "francês"
const prefixHelloSpanish = "Hola, "
const prefixHelloFrench = "Bonjour, "
const prefixHelloPortuguese = "Olá, "

func Ola(name string, language string) string {
	if name == "" {
		name = "mundo!"
	}

	return prefixGreeting(language) + name
}

func prefixGreeting(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = prefixHelloSpanish
	case french:
		prefix = prefixHelloFrench
	default:
		prefix = prefixHelloPortuguese
	}
	return
}

func main() {
	fmt.Println(Ola("mundo", ""))
}
