package main

import (
	"fmt"
	"os"

	"github.com/Ladence/golang_base_rest/internal/server"
)

func main() {
	if err := server.Run(8080); err != nil {
		fmt.Println("Error on server running")
		os.Exit(1)
	}
}
