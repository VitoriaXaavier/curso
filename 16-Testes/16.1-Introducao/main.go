package main

import (
	"fmt"
	"introducao-testes/endereco"
)
func main() {
	tipoEndereco := endereco.TipoEndereco("Avenida Rio branco")

	fmt.Println(tipoEndereco)
}