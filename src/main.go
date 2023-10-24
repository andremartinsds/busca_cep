package main

import (
	"busca-cep/src/resources"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	//fileServer := http.FileServer(http.Dir("./src/public"))

	//mux.Handle("/", fileServer)

	mux.HandleFunc("/cep", resources.BuscaCepResource)

	log.Print("starting server")
	err := http.ListenAndServe(":8081", mux)

	log.Fatal(err)
}
