package main

import (
	"booking/cmd/server/config"
	"booking/cmd/server/database"
	"fmt"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	database.NewDB()
	s := config.NewGRPCServer()

	fmt.Println("Server listening on port 8080")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
