package hello

import "fmt"

const prefixHelloPortuguese = "Olá, "

func Ola(nome string) string {
	if nome == "" {
		nome = "mundo!"
	}

	return prefixHelloPortuguese + nome
}

func main() {
	fmt.Println(Ola("mundo"))
}
