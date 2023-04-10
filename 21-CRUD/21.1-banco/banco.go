package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConectaBanco() (*sql.DB, error) {
	stringConexao := ("root:Viviayla21*@tcp(localhost:3306)/devbook?parseTime=true&loc=Local")

	db, err := sql.Open("mysql", stringConexao)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
