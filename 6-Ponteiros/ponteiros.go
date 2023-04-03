package main

import "fmt"

func main() {
	var variavel int
	var ponteiro *int

	variavel = 100
	ponteiro = &variavel

	fmt.Println(variavel, ponteiro) // mostra o endereço de memória
	fmt.Println(variavel, *ponteiro) //pega o valor que está no endereço de memória - desreferenciação

	variavel = 150

	fmt.Println(variavel, ponteiro) // o endereço de memória é o mesmo, apesar do valor ser diferente
	fmt.Println(variavel, *ponteiro)

}