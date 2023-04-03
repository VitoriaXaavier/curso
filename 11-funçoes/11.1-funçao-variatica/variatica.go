package main

import "fmt"

func soma(numeros ...int) int {

	total := 0

	for _, numero := range numeros {
		total += numero
	}
	return total
}

func escrever(texto *string, numeros ...int) {
	*texto = "Ola Vitoria"
	for _, numero := range numeros {
		fmt.Println(*texto, numero)

	}
}

func main() {
	somaTotal := soma(1, 2, 3, 3, 4)
	fmt.Println(somaTotal)
	saudacao := "Ola mundo"
	println(saudacao)
	escrever(&saudacao, 1, 2, 3, 4, 5, 6)
	println(saudacao)

}
