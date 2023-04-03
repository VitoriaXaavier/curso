package main

import "fmt"

func main() {

	func () {
		fmt.Println("Função anonima simples")
	}()
	

	func (texto string)  {
		fmt.Println(texto)
	}("Função anonima com parametro")



	retorno := func (texto string) string  {

	return	fmt.Sprintf("Função anonima com parametro %s", texto)
		
	}("e retorno")

	fmt.Println(retorno)

}