package main

import (
	"GREEN-API/config"
	"GREEN-API/server"
	"fmt"
	"log"
)

func main() {
	cfg := config.Load()

	srv := server.New(cfg)

	addr := fmt.Sprintf(":%d", cfg.Port)
	log.Printf("Starting server on %s", addr)

	if err := srv.Start(addr); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
