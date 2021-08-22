package main

import (
	"fmt"
)

func main() {
	canal := make(chan string, 2) //com buffer só vai bloquear quando atingir a capacidade máxima

	canal <- "Olá Mundo"
	canal <- "Programando em Go!" //dois valores, pois foi setado '2' como limite

	mensagem := <- canal
	mensagem2 := <- canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}
