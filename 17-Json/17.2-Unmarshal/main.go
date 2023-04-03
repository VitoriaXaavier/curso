package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome string `json: "nome"`
	Raca string `json: "raca"`
	Idade uint `json: "idade"`
}

func main() {

	cachorroEmJSON := `{"nome": "Ayla", "raca": "bulldog","idade":1}`

	// imprime o JSON
	fmt.Println(cachorroEmJSON)

	var c cachorro

	// converte o json em um struct e trata o erro
	err := json.Unmarshal([]byte(cachorroEmJSON), &c)
		if err != nil {
			log.Fatal(err)
		}

	// imprime a struct
 	fmt.Println(c)

	//MAP

	cachorroMapEmJson := `{"nome": "Jacob", "raca": "Yorkshire"}`

	c2 := make(map[string]string)

		err = json.Unmarshal([]byte(cachorroMapEmJson), &c2)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(c2)
}