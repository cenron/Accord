package main

import (
	"accord/server/db"
	"accord/server/internal/user"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	fmt.Println("Accord chat...")

	db, err := db.NewMongoDatabase()
	if err != nil {
		log.Fatalf("could not initalize database connection: %s\n", err)
	}

	userHandler := user.NewUserHandler(db)

	r := gin.Default()
	r.POST("/test", userHandler.CreateUser)

	r.Run(":8000")
}
