package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/akwanmaroso/nakama-app/handler"
	"github.com/akwanmaroso/nakama-app/internal/service"
	"github.com/hako/branca"
	_ "github.com/jackc/pgx/stdlib"
)

const (
	databaseURL = "postgresql://root@127.0.0.1:26257/nakama?sslmode=disable"
	port        = 3000
)

func main() {
	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		log.Fatalf("could not open db connection: %v\n", err)
		return
	}

	defer db.Close()
	if err = db.Ping(); err != nil {
		log.Fatalf("could not ping to db : %v\n", err)
		return
	}

	codec := branca.NewBranca("supersecretkeyyoushouldnotcommit")
	codec.SetTTL(uint32(service.TokenLifespan.Seconds()))
	s := service.New(db, codec)

	h := handler.New(s)
	addr := fmt.Sprintf(":%d", port)
	log.Printf("accepting connections on port %d", port)
	if err = http.ListenAndServe(addr, h); err != nil {
		log.Fatalf("could not start server: %v\n", err)
	}
}
