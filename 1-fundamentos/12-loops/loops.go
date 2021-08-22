package main

import (
	//"time"
	"fmt"
)

func main() {
	//i := 0

	//for i < 5 {
		//i++
		//fmt.Println("incrementando i")
		//time.Sleep(time.Second)

	//}

	//fmt.Println(i)

	//for j :=0; j < 4; j++ {
		//fmt.Println("incrementando j: ", j)
		//time.Sleep(time.Second)
	//}


	nomes := []string{"Júlio", "Daniela", "Laura"}

	for indice, nome := range nomes {
		fmt.Println(indice, " => ", nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, " => ", string(letra))
	}

	//OBS: não é possível utilizar range em struct.
}