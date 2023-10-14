package main

import (
	"busca-cep/src/resources"
	"net/http"
)

func main() {
	http.HandleFunc("/", resources.BuscaCepResource)
	http.ListenAndServe(":8081", nil)
}
