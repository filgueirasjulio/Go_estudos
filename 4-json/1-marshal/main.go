package main

import (
	"bytes"
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
	g := Gato{"Tupy", "SRN", 1}
	
	gatoEmJSON, erro := json.Marshal(g)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(gatoEmJSON)
	fmt.Println(bytes.NewBuffer(gatoEmJSON))
}