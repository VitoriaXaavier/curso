package metodo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type usuario struct {
	ID    uint64 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Erro ao ler o corpo da requisição"))
		return
	}

	var usuario usuario

	if err = json.Unmarshal(corpoRequisicao, &usuario); err != nil {
		w.Write([]byte("Erro ao converter o usuario para struct"))
		return
	}
	fmt.Println(usuario)
}
