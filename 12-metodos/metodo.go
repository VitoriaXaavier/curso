package main

import "fmt"

type usuario struct{
	nome string
	idade uint
}

func (u usuario) salvar() {
	fmt.Printf("Salvando o usuario %s no banco de dados \n", u.nome)
}

func( u usuario) maiorIdade() bool{
	return u.idade >= 18
}

func (u *usuario) aniversario() {
	u.idade++
}

func main(){

	usuario1 := usuario{"Vitoria",23}

	fmt.Println(usuario1)
	usuario1.salvar()

	maiorIdade := usuario1.maiorIdade()
	fmt.Println("Maior de idade? ",maiorIdade)

	fmt.Println("Fazendo aniversário")
	usuario1.aniversario()
	fmt.Println(usuario1.idade)
}