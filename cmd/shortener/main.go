package main

import (
	"fmt"
	"net/http"

	"github.com/IndianaDP/net-http/internal/config"
	"github.com/IndianaDP/net-http/internal/routers"
)

func main() {

	cfg, err := config.LoadConfig(true)
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	routers.SetupRouter()

	fmt.Println("Starting server at port", cfg.Address)
	if err := http.ListenAndServe(cfg.Address, nil); err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
