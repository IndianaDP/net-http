package main

import (
	"fmt"
	"net/http"

	"github.com/IndianaDP/net-http/internal/app/config"
	"github.com/IndianaDP/net-http/internal/app/routers"
)

func main() {
	cfg, err := config.LoadConfig(true)
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	handler := routers.SetupRouter(cfg)

	fmt.Println("Starting server at port", cfg.Address)
	if err := http.ListenAndServe(cfg.Address, handler); err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
