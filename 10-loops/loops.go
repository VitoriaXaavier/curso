package main

import (
	"fmt"
	"time"
)

func main() {
	i := 10

	for i<  15 {
		i++
	
	}

	fmt.Println("primeiro",i)

	for j := 1; j < 10; j++ {
		time.Sleep(time.Second)
		fmt.Println("segundo",j)
	}

	nomes := [3]string{"Vitoria","Andre","Ayla"}

	for indice, valor := range nomes{
		fmt.Println("array",indice,valor)
	}

	for indice, letra := range "PALAVRA"{
		fmt.Println("range",indice,string(letra))
	}

	usuario := map[string] string {
		"nome": "Vitoria",
		"sobrenome": "Alves",
	}

	for chave, valor := range usuario{
		fmt.Println("map",chave,valor)
	}


}