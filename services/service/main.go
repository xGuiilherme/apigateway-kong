package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// Aqui ele vai mostrar o resultado/conteudo que tambem esta em uma env de ambiente.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, os.Getenv("CONTENT"))
}

// Quando alguem acessar minha url/service baseado em uma porta que vai ser uma env de ambiente.
func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), nil))
}
