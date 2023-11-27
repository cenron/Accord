package router

import (
	"accord/internal/ws"
	"accord/pkg/db"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"time"
)

func InitRouter(store *db.MongoStore, log *log.Logger) chi.Router {

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/api", func(r chi.Router) {

		//userRouter(r, store)
		webSocketRouter(r, store)

	})

	return r
}

func webSocketRouter(rg chi.Router, store *db.MongoStore) {
	wsHandler := ws.NewWsHandler(store)

	rg.Route("/v1/ws", func(r chi.Router) {
		r.Get("/", wsHandler.JoinRoom)
	})
}

/*func userRouter(rg chi.Router, store *db.MongoStore) {
	userHandler := user.NewUserHandler(store)

	rg.Route("/v1/user", func(r chi.Router) {
		r.Post("/signup", userHandler.Signup)
		r.Post("/login", userHandler.Login)
		r.Post("/logout", userHandler.Logout)
	})
}*/
