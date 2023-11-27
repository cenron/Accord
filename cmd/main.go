package main

import (
	"accord/internal/router"
	"accord/pkg/db"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
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

	// Create a route along /files that will serve contents from
	// the ./data/ folder.
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "web"))
	FileServer(r, "/", filesDir)

	if err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", ServerPort), r); err != nil {
		logger.Panicf("could not start server on port: %d", ServerPort)
	}
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
