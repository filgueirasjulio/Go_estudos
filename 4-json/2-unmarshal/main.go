package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Gato struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	gatoEmJSON := `{"nome":"Tupy", "raca":"SRN", "idade":1}`

	var g Gato 

	if erro:= json.Unmarshal([]byte(gatoEmJSON), &g); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(g)
}