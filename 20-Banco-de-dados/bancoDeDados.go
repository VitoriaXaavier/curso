package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	stringConexao := ("root:Viviayla21*@tcp(localhost:3306)/devbook?parseTime=true&loc=Local")
	db, err := sql.Open("mysql", stringConexao)

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// testa a conexão
	if err = db.Ping(); err != nil {
		panic(err.Error())
	}
	fmt.Println("conexao aberta")
	linhas, err := db.Query("select * from usuarios")

	if err != nil {
		panic(err.Error())
	}
	defer linhas.Close()
	fmt.Println(linhas)

}
