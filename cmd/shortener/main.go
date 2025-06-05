package main

import (
	"fmt"
	"net/http"

	"github.com/IndianaDP/net-http/internal/routers"
)

func main() {
	routers.SetupRouter()

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
