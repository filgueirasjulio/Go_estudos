package main

import "fmt"

func main() {
	var array1 [5]int
	array1[0] = 1
	fmt.Println(array1)

	array2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4 ,5}
	fmt.Println(array3)

	//todos exemplos acima são dinamicos, não é possivel adicionar um novo indices posteriormente

	//com slices podemos fazer isso
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	slice = append(slice, 6)
	fmt.Println(slice)

	slice2 := array2[1:3] //o primeiro indíce é inclusivo e o segundo exclusivo
	fmt.Println(slice2)
}