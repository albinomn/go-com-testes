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

	prefix := prefixHelloPortuguese

	if language == spanish {
		prefix = prefixHelloSpanish
	}

	if language == french {
		prefix = prefixHelloFrench
	}

	return prefix + name
}

func main() {
	fmt.Println(Ola("mundo", ""))
}
