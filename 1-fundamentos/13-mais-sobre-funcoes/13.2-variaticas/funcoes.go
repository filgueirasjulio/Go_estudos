package main

import "fmt"

func escrever(texto string, numeros ...int) {
	total := 0

	for _, numero := range numeros {
		total += numero
	}
	
	fmt.Println("O ", texto, " foi ", total)
}

func main() {
	escrever("somátorio", 10, 20, 12, 45, 32, 56, 28, 110, 200)
}