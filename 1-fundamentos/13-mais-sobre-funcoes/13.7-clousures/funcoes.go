package main

import "fmt"

func clousure() func() {
	texto := "Dentro da funcão clousure"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}

func main() {
	texto := "Dentro da função main"
	fmt.Println(texto)

	funcaoNova := clousure()
	funcaoNova()
}