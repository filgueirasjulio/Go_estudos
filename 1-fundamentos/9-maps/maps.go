package main

import "fmt"

func main() {
	usuario := map[string]string{
		"nome":	"Júlio",
		"profissao": "Programador",
	}

	fmt.Println(usuario)

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Júlio",
			"sobrenome": "Filgueiras",
		},
		"profissao": {
			"cargo": "Programador",
			"senioridade": "Jr",
		},
	}

	fmt.Println(usuario2)
	//deletar chave
	delete(usuario2, "nome")
	fmt.Println(usuario2)
	//acrescentando chave
	usuario2["naturalidade"] = map[string]string{
		"cidade": "Salvador",
	}
	fmt.Println(usuario2)
}