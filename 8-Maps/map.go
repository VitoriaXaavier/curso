package main

import "fmt"

func main () {
	usuario := map[string]string {
		"nome": "Vitoria",
		"sobrenome": "Alves",
	}

	fmt.Println(usuario["nome"])
	fmt.Println(usuario)


	// map aninhado

	usuario2 := map[string]map[string]string {

		"nome": {
			"first": "Vitoria",
			"last": "Alves",
		},
		"profissao": {
			"formacao": "desenvolvedora",
			"empresa": "globo",
		},
	}

	fmt.Println(usuario2)
	fmt.Println(usuario2["nome"])

	// deletar uma chave

	delete(usuario2,"nome")

	//adicionar uma chave

	usuario2["endereco"] = map[string]string{
		"cidade": "araruama",
		"estado": "rj",
	}

	fmt.Println(usuario2)
}