package main

import "fmt"

type pessoa struct {
	nome  string
	idade uint8
	cpf   int
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {

	pessoa1 := pessoa{"Joao",20,1231231231}
	fmt.Println(pessoa1)

	estudante1 := estudante{pessoa1,"engenharia","usp"}
	fmt.Println(estudante1)
	fmt.Println(estudante1.nome,estudante1.curso)
}
