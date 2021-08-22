package main

import "fmt"

func somar(n1, n2 int8) (int8) {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	somar := n1 + n2
	subtrair := n1 -n2
	return somar, subtrair
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	f := func(txt string) (string) {
		return txt
	}

	resultadof := f("Texto função f")
	fmt.Println(resultadof)

	resSoma, resSub := calculosMatematicos(20, 10)
	fmt.Println(resSoma, resSub)
}