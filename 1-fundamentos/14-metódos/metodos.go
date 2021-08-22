package main

import "fmt"

type usuario struct {
	nome string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando usuário %s no banco...\n", u.nome);
}

func (u usuario) maiorDeIdade() {

  if u.idade >= 18 {
	  fmt.Println("maior de idade")
  } else {
	  fmt.Println("menor de idade")
  }
} 

func (u *usuario) fazerAniversario() {
	u.idade ++
}

func main() {
	usuario1 := usuario{"Júlio", 33}
	fmt.Println(usuario1)
	usuario1.salvar()
	usuario1.maiorDeIdade()
	usuario1.fazerAniversario()
	fmt.Println(usuario1.idade);
}