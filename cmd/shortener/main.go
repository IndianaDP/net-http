package main

import (
	"fmt"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/IndianaDP/net-http/internal/app/config"
	"github.com/IndianaDP/net-http/internal/app/db/storage"

	"github.com/IndianaDP/net-http/internal/app/routers"
)

func main() {
	cfg, err := config.LoadConfig(true)
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	var handler http.Handler
	storage, err := storage.NewStorage(cfg)
	if err != nil {
		fmt.Println("Error init storage:", err)
		return
	}
	handler = routers.SetupRouter(cfg, storage)

	fmt.Println("Starting server at port", cfg.Address)
	if err := http.ListenAndServe(cfg.Address, handler); err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
