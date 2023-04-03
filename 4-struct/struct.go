package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
	endereco endereco
}
type endereco struct {
	logradouro string
	numero uint8
}
func main() {

	var u usuario
	fmt.Println(u)
	u.nome = "Vitoria"
	u.idade = 23
	fmt.Println(u.nome, u.idade)

enderecoEx := endereco{"Rua dos bobos",0}
	usuario2 := usuario{nome: "Vitoria Alves",idade: 23}

	fmt.Println(usuario2)

	usuario3 := usuario{"Andre", 35, enderecoEx}
	fmt.Println(usuario3)

	usuario4 := usuario{nome: "Ayla"}
	fmt.Println(usuario4)

}
