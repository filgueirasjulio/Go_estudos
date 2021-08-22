package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
		case 1:
			return "Domingo"
		case 2:
			return "Segunda-feira"
		case 3:
			return "Terça-feira"
		case 4:
			return "Quarta-feira"
		case 5:
			return "Quinta-feira"
		case 6:
			return "Sexta-feira"
		case 7:
			return "Sábado"
		default:
			return "Dia desconhecido"
	}
}

func diaDaSemana2 (numero int) string {
	switch {
		case numero == 1:
			return "Domingo"
		case numero == 2:
			return "Segunda-feira"
		case numero == 3:
			return "Terça-feira"
		case numero == 4:
			return "Quarta-feira"
		case numero == 5:
			return "Quinta-feira"
		case numero == 6:
			return "Sexta-feira"
		case numero == 7:
			return "Sábado"
		default:
			return "Dia desconhecido"
		}
}

func diaDaSemana3 (numero int) string {
	var diaDaSemana3 string

	switch {
		case numero == 1:
			diaDaSemana3 = "Domingo"
		case numero == 2:
			diaDaSemana3 = "Segunda-feira"
		case numero == 3:
			diaDaSemana3 = "Terça-feira"
		case numero == 4:
			diaDaSemana3 = "Quarta-feira"
		case numero == 5:
			diaDaSemana3 = "Quinta-feira"
		case numero == 6:
			diaDaSemana3 = "Sexta-feira"
		case numero == 7:
			diaDaSemana3 = "Sábado"
		default:
			diaDaSemana3 = "Dia desconhecido"
		}

	return diaDaSemana3
}

func main() {
	dia := diaDaSemana(5)
	dia2 := diaDaSemana2(6)
	dia3 := diaDaSemana3(1)
	fmt.Println(dia)
	fmt.Println(dia2)
	fmt.Println(dia3)
}