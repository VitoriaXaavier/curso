package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConectaBanco() (*sql.DB, error) {
	// sintaxe básica: usuario:senha@tcp(endereco_servidor:porta)/banco_de_dados.
	stringConexao := ("root:Viviayla21*@tcp(localhost:3306)/devbook?parseTime=true&loc=Local")
	//parseTime: deve fazer a  conversão automática dos valores de data e hora retornados pelo banco de dados para o
	//formato da linguagem Go
	//loc=Local informa ao driver MySQL do Go que as informações de data e hora devem ser interpretadas utilizando o
	//fuso horário local da máquina onde o código está sendo executado.

	db, err := sql.Open("mysql", stringConexao)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
