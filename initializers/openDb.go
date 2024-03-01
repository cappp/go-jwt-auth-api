package initializers

import (
	"database/sql"
	"log"
	"os"
)

var DB *sql.DB

func OpenDb() {
	os.Remove("./users.db")
	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		log.Fatal("Erro ao abrir o banco de dados:\n" + err.Error())
	}
	_, err = db.Exec("create table users (name text, username text, password text)")
	if err != nil {
		log.Fatal("Erro ao criar a tabela 'users':\n" + err.Error())
	}
	DB = db
}
