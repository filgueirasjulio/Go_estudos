package main

import "fmt"

func fibonacci(posicao uint) uint {
	if posicao  <= 1 {
		return posicao;
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1);
}

func main() {

	posicao := uint(50)

	for x := uint(0); x <= posicao; x++ {
		fmt.Println(fibonacci(x))
	}  
	
}