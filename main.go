package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"techschool/api"
	db "techschool/db/sqlc"
)

const (
	dbDriver   = "postgres"
	dbSource   = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	serverAddr = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddr)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
