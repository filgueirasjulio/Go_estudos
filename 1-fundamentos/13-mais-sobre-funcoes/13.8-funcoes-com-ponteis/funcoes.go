package main

import "fmt"

func inverterSinal(numero int) int {
	numero = numero * -1
	return numero
}

func inverterSinalComPonteiro(numero *int) int {
	*numero = *numero * -1
	return *numero
}

func main() {
	numero := 20
	fmt.Println(numero)
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido)

	novoNumero := 40
	fmt.Println(novoNumero)
	inverterSinalComPonteiro(&novoNumero)
	fmt.Println(novoNumero)
}
