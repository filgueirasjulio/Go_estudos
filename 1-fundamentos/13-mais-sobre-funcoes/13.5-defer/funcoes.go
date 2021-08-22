package main

import (
	"fmt"
)

func alunoEstaAprovado(n1, n2 float32) bool {
	media := (n1 + n2)/2

	defer fmt.Println("Média calculada. O resultado será retornado")
	fmt.Println("Entrando na função para verificar se o aluno foi aprovado")

	if (media >= 6) {
		return true;
	}

	return false;
}

func main() {
	fmt.Println(alunoEstaAprovado(7, 8))
}