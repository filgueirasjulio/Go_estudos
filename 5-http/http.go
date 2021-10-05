package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Você está na home"))
	})

	log.Fatal(http.ListenAndServe(":5000", nil))
}