package main

import (
	"github.com/jonattasmoraes/app-go/internal/config"
	"github.com/jonattasmoraes/app-go/internal/server"
)

func main() {
	// Initialize config
	config.InitializeConfig()

	// Initialize server
	server.InitializeRouter()
}
