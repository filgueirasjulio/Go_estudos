package main

import (
	"fmt"
	"go-estudo/3-testes_automatizados/1-introducao/enderecos"
)

func main() {
	TipoDeEndereco := enderecos.TipoDeEndereco("Avenida 7 de Setembro")
	fmt.Println(TipoDeEndereco)
}