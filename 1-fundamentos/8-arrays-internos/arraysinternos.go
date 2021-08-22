package main

import "fmt"

func main() {
	//quando se cria slice sem referência direta a um array
	slice := make([]float32, 10, 11)
	fmt.Println(slice)
	fmt.Println(len(slice)) //length
	fmt.Println(cap(slice))//capacidade
	slice = append(slice, 1)
	slice = append(slice, 2)
	fmt.Println(slice)
	fmt.Println(len(slice)) //length
	fmt.Println(cap(slice))//capacidade

	slice2 := make([]float32, 10)
	slice2 = append(slice, 11)
	slice2 = append(slice, 12)
	fmt.Println(slice2)
	fmt.Println(len(slice2)) //length
	fmt.Println(cap(slice2))//capacidade
}