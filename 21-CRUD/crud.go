package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/VitoriaXaavier/curso/21-CRUD/metodo"

	"github.com/gorilla/mux"
)

func main() {
	//cria um roteador vazio
	router := mux.NewRouter()
	//registro das rotas e func handlers
	router.HandleFunc("/usuario", metodo.CriarUsuario)

	fmt.Println("escutando na porta 5000")
	//inicializa o roteador http
	log.Fatal(http.ListenAndServe(":5000", router))
}
