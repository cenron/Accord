package main

import (
	"accord/internal/user"
	"accord/pkg/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	fmt.Println("Accord chat...")

	client, err := db.NewMongoDatabase()
	if err != nil {
		log.Fatalf("could not initalize database connection: %s\n", err)
	}

	defer client.Close()

	userHandler := user.NewUserHandler(client)

	r := gin.Default()
	r.POST("/test", userHandler.CreateUser)

	r.Run(":8000")
}
