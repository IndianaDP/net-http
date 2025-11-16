package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/IndianaDP/net-http/internal/app/config"
	"github.com/IndianaDP/net-http/internal/app/db"
	"github.com/IndianaDP/net-http/internal/app/routers"
)

func main() {
	connStr := "postgres://postgres:secret@localhost:5433/db-store?sslmode=disable"

	conn, err := db.NewDB(connStr)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer conn.Close()

	cfg, err := config.LoadConfig(true)
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	handler := routers.SetupRouter(cfg, conn)

	fmt.Println("Starting server at port", cfg.Address)
	if err := http.ListenAndServe(cfg.Address, handler); err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
