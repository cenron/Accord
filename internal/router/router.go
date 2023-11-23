package router

import (
	"accord/internal/user"
	"accord/pkg/db"
	"github.com/gin-gonic/gin"
	"log"
)

func InitRouter(store *db.MongoStore, log *log.Logger) *gin.Engine {

	r := gin.Default()
	routeGroup := r.Group("/api")

	// Set up our user routes.
	userRouter(routeGroup, store)

	return r
}

func userRouter(rg *gin.RouterGroup, store *db.MongoStore) {
	userHandler := user.NewUserHandler(store)

	v1 := rg.Group("/v1/user")
	{
		v1.POST("/signup", userHandler.Signup)
		v1.POST("/login", userHandler.Login)
		v1.POST("/logout", userHandler.Logout)
	}
}
