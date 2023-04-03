package main

import (
	"fmt"
	"reflect"
)

func main() {
	var array [5] string
	array[0] = "Posição 1, indice 0"
	fmt.Println(array)

	array2 := [5] string{ "Posição 1, indice 0", "Posição 2, indice 1","Posição 3, indice 2","Posição 4, indice 3","Posição 5, indice 4"}
	fmt.Println(array2)

	// array com constante

	type Currency int

	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)

	symbol := [...] string{ USD: "$", EUR: "€", GBP: "£", RMB: "¥"}

	fmt.Println("Posição", RMB, "Moeda", symbol[RMB])
	fmt.Println("Posição", USD, "Moeda",symbol[USD],"Posição", EUR, 
	"Moeda",symbol[EUR], "Posição", GBP, "Moeda", symbol[GBP], "Posição", RMB,"Moeda",symbol[RMB])

	//infere o tamanho do array pela quantidade de posições
	array3 := [...]int{1,2,3,4,5}
	fmt.Println(array3)

	// slice

	slice := []int{1,2,3,4,5}
	fmt.Println(slice)

	// mostra o tipo
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array2))

	slice = append(slice, 11,13) // devolve o slice com o item adicionado
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	
	slice2 = append(slice2, "posição apepend")
	fmt.Println(slice2)
	
	array2[2] = "Array alterado"
	fmt.Println(slice2)

	

}