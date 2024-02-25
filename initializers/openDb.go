package initializers

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func OpenDb() {
	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		log.Fatal("Erro ao abrir o banco de dados:\n" + err.Error())
	}
	DB = db
}
