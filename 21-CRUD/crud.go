package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/VitoriaXaavier/curso/metodo"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/usuario", metodo.CriarUsuario)

	fmt.Println("escutando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
