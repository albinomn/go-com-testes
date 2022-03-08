package hello

import "fmt"

func Ola(nome string) string {
	return "Olá, " + nome + "!"
}

func main() {
	fmt.Println(Ola("mundo"))
}
