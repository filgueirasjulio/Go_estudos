package main

import "fmt"

type pessoa struct {
	nome string
	idade uint8
	altura uint8
}

type estudante struct {
	pessoa
	curso string
	faculdade string
}

func main() {
	p1 := pessoa{"Júlio", 33, 180}
	fmt.Println(p1)

	e1 := estudante{p1, "Sistemas de Informação", "Unifacs"}
	fmt.Println(e1)

	//para acessar uma proprieda, pode fazer de forma direta
	fmt.Println(e1.nome)
	//não é necessário fazer e1.pessoa.nome
}