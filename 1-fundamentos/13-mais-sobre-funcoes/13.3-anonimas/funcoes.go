package main

import "fmt"

func main() {
	resultado := func(texto string) string {
		return fmt.Sprintf("Passando -> resultado %s", texto)
	}("texto como parâmetro")

	fmt.Println(resultado)
}