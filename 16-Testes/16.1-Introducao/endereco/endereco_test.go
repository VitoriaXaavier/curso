package endereco

import "testing"

type cenarioTest struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestEndereco(t *testing.T) {
	
	cenariosTest := []cenarioTest {
		
		{"Rua abc", "Rua"},
		{"Avenida Rio Branco", "Avenida"},
		{"Estrada do Pacifico", "Estrada"},
		{"Praça da Bandeira", "Não é um tipo válido"},
		{"RUA ABC", "Rua"},
		{" ", "Não é um tipo válido"},
		{"avenida rio branco", "Avenida"},
	}
	
	for _,cenario := range cenariosTest {
		
		tipoEnderecoRecebido := TipoEndereco(cenario.enderecoInserido)
		if tipoEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O endereço esperado era %s, mas foi recebido %s", 
			cenario.retornoEsperado, 
			tipoEnderecoRecebido)
		}
	}
}
//Mostra a % de cobertura do teste: go test --cover 

//Para alocar a cobertura em um arquivo separado: go test --coverprofile cobertura.txt

//Para ler o arquivo que foi gerado: go tool cover --func=cobertura.txt

//Para saber qual parte do codigo não está sendo coberta pelo teste: 
// go tool cover --html=cobertura.txt - Vai gerar uma página html temporaria que irá mostrar 
// as partes que estão sendo testadas e as partes que não estão sendo testadas

