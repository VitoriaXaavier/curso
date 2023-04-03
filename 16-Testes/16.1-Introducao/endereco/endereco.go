package endereco

import (
	"strings"
)


func TipoEndereco(endereco string) string {
	tiposValidos := []string {"rua", "avenida", "rodovia", "estrada"}

	enderecoMinusculo := strings.ToLower(endereco)
	primeiraPalavra := strings.Split(enderecoMinusculo, " ")[0]
	enderecoTemUmTipoValido := false


	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavra {
			enderecoTemUmTipoValido = true
		}
	}
	if enderecoTemUmTipoValido {
		return strings.Title(primeiraPalavra)
	} 
	return "Não é um tipo válido"
}