package metodo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/VitoriaXaavier/curso/21-CRUD/banco"
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

	db, err := banco.ConectaBanco()
	if err != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
	}

	statement, err := db.Prepare("insert into usuarios (nome, email) values (?,?)")

	if err != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}

	defer statement.Close()

	//executa uma instruçao que foi preparada com o metodo db.Prepare
	insercao, err := statement.Exec(usuario.Nome, usuario.Email)
	if err != nil {
		w.Write([]byte("Erro ao executar o statement"))
		return
	}

	idInserido, err := insercao.LastInsertId() // retorna o valor da  chave primaria
	if err != nil {
		w.Write([]byte("Erro ao devolver o ID"))
		return
	}

	//status code

	w.WriteHeader(http.StatusCreated)

	w.Write([]byte(fmt.Sprintf("Usuario inserido com suscesso, ID: %d", idInserido)))
	fmt.Println(idInserido,usuario.Nome, usuario.Email)
}
