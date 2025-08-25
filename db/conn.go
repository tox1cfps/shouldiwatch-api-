package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectToDB(HOST, PORT, USER, PASSWORD, DBNAME, SSLMODE string) (*sql.DB, error) {
	connStr := "host=" + HOST + " port=" + PORT + " user=" + USER + " password=" + PASSWORD + " dbname=" + DBNAME + " sslmode=" + SSLMODE
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("erro ao conectar ao banco:", err)
	}

	log.Println("Conectado com sucesso ao banco de dados!")

	DB = db

	return db, nil
}
