package main

import (
	"errors"
	"fmt"
)

func main() {
	//existe o int8, int16, int32, int64
	//quanto maior o numero, maior  a capacidade de armazenamento
	//se usar apenas int, ele sera 32 ou 64 de acordo com a arquitetura da sua máquina
	var numero int16 = 1000
	fmt.Println(numero)

	//var numero2 uint16 = -1000  //o uint impede números negativos

	//alias
	var numero3 rune = 1234  //rune é o mesmo que int32
	var numero4 byte = 123 //é o mesmo que uint8
	fmt.Println(numero3, numero4)

	var numero5 float32 = 12.3234
	var numero6 float64 = 14.5323
	fmt.Println(numero5, numero6)

	var str string = "Texto"
	fmt.Println(str)

	char := 'A' //não existe char, ele converte para int indicando o valor na tabela ASCII
	fmt.Println(char)

	var booleano = true
	fmt.Println(booleano)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}