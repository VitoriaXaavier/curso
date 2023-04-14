package metodo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/VitoriaXaavier/curso/21-CRUD/banco"

	"github.com/gorilla/mux"
)

type usuario struct {
	ID    uint64 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	// ler todos os dados de um objeto io.Reader e retorna
	// um slice de bytes ([]byte) contendo esses dados. No caso os dados digitados no body
	corpoRequisicao, err := ioutil.ReadAll(r.Body)
	if err != nil {
		//escreve dados na resposta HTTP que será enviada ao cliente.
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

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, err := banco.ConectaBanco()
	if err != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}
	defer db.Close()
	linhas, err := db.Query("select * from usuarios")
	if err != nil {
		w.Write([]byte("Erro ao selecionar as linhas da tabela"))
		return
	}
	defer linhas.Close()

	var usuarios []usuario

	for linhas.Next() {
		var usuario usuario

		if err := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); err != nil{
			w.Write([]byte("Erro ao scannear o usuario"))
			return
		}
		usuarios = append(usuarios, usuario)
	}
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(usuarios); err != nil{
		w.Write([]byte("Erro ao converter os usuarios para json"))
		return
	}
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)

	ID, err :=  strconv.ParseUint(parametros["id"], 10, 32)

	if err != nil {
		w.Write([]byte ("Erro ao converter a string para inteiro"))
		return
	}

	db, err := banco.ConectaBanco()
	if  err != nil {
		w.Write([]byte("Erro ao conectar com o banco  de dados"))
		return
	}
	defer db.Close()
	// foi usado a Query, pois aqui iremos retornar resultados diferentes cada vez que executarmos 
	linha, err := db.Query(" select * from usuarios where  id = ? ",  ID)
	if err != nil {
		w.Write([]byte("Erro ao selecionar o id"))
		return
	}
	defer linha.Close()
	var usuario usuario

	if linha.Next() {
		if err := linha.Scan(&usuario.ID,&usuario.Nome,&usuario.Email); err != nil{
			w.Write([]byte("Erro ao scannear o usuario"))
			return
		}

		if err  := json.NewEncoder(w).Encode(usuario); err != nil {
			w.Write([]byte("Erro ao transformar em json"))
			return
		}
	}
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	ID, err := strconv.ParseUint(parametros["id"], 10, 32)

	if err  != nil {
		w.Write([]byte("Erro ao converter a string id para uint id"))
		return
	}
	corpoRequisicao, err := ioutil.ReadAll(r.Body)
	if err != nil  {
		w.Write([]byte("Erro ao ler o corpo da requisiçao"))
		return
	}
	var usuario usuario
	if err := json.Unmarshal(corpoRequisicao,&usuario); err != nil {
		w.Write([]byte("Erro ao converter o json"))
		return
	}

	db, err := banco.ConectaBanco()
	if err != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}
	defer db.Close()

	statemant, err := db.Prepare("update usuarios set nome = ?, email = ? where id = ?")
	if err != nil {
		w.Write([]byte("Erro ao atualizar usuario no banco"))
		return
	}
	defer statemant.Close()
	if _, err := statemant.Exec(usuario.Nome, usuario.Email, ID); err != nil {
		w.Write([]byte("Erro ao atualizar usuario"))
		return
	}
w.WriteHeader(http.StatusNoContent)
}

func DeletaUsuario (w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	ID, err := strconv.ParseUint(parametros["id"], 10, 32)

	if err  != nil {
		w.Write([]byte("Erro ao converter a string id para uint id"))
		return
	}
	db, err := banco.ConectaBanco()
	if err != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}
	defer db.Close()
	statement, err := db.Prepare("delete from usuarios where id = ?")
	if err != nil {
		w.Write([]byte("Erro ao preparar o statement"))
		return
	}
	defer statement.Close()
	if _, err := statement.Exec(ID); err != nil{
		w.Write([]byte("Erro ao deletar o usuario"))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}