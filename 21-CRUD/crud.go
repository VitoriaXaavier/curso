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
	router.HandleFunc("/usuarios", metodo.CriarUsuario).Methods(http.MethodPost)
	router.HandleFunc("/usuarios", metodo.BuscarUsuarios).Methods(http.MethodGet)
	router.HandleFunc("/usuario/{id}", metodo.BuscarUsuario).Methods(http.MethodGet)


	fmt.Println("escutando na porta 5000")
	//inicializa o roteador http
	log.Fatal(http.ListenAndServe(":5000", router))
}
