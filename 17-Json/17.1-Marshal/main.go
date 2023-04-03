package main

import (
	"bytes"
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
	c := cachorro{"Ayla","bulldog frances", 5}

	fmt.Println(c)  // devolve a struct 

	cachorroJSON, err := json.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cachorroJSON) // irá devolver um slice de bytes

	// devolve o arquivo em json
	fmt.Println(bytes.NewBuffer(cachorroJSON)) // irá converter o slice de bytes para nosso entendimento 

	fmt.Println("------------------------------------------------------------------------------------")

	// MAP

	c2 := map[string] string{
		"nome": "Jacob",
		"raca": "Yorkshire",
	}

	// imprime um tipo map
	fmt.Println(c2)

	// converte para json
	cachorroMapJSON, err := json.Marshal(c2)

	// imprime um slice de bytes
	fmt.Println(cachorroMapJSON)

	// imprime o json, convertendo o slice de bytes em algo legível
	fmt.Println(bytes.NewBuffer(cachorroMapJSON))

}