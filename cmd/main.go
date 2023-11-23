package main

import (
	"accord/internal/router"
	"accord/pkg/db"
	"fmt"
	"log"
)

const ServerPort int32 = 8080

func main() {

	logger := log.Default()

	client, err := db.NewMongoDatabase()
	if err != nil {
		log.Fatalf("could not initalize database connection: %s\n", err)
	}

	defer client.Close()

	r := router.InitRouter(client, logger)

	if err := r.Run(fmt.Sprintf("0.0.0.0:%d", ServerPort)); err != nil {
		logger.Panicf("could not start server on port: %d", ServerPort)
	}
}
