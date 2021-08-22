package main

import "fmt"

type usuario struct {
	nome string
	idade uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero uint8
}

func main() {
	var u usuario
	u.nome = "Júlio"
	u.idade = 33
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua da bica", 2}

	u2 := usuario{"Roberto", 68, enderecoExemplo}
	fmt.Println(u2)

	//se quiser passar apenas um valor
	u3 := usuario{nome: "Mariana"}
	fmt.Println(u3)
}