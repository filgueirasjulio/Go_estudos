package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"crud/servidor"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/usuarios", servidor.CriarUsuario).Methods("POST")

	fmt.Println("Escutando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}