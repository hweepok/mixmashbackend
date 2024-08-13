package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/hweepok/mixmashbackend/pkg/service/recipe"
	"github.com/hweepok/mixmashbackend/pkg/service/user"
	"github.com/hweepok/mixmashbackend/pkg/storage"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()

	userStore := db.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(router)

	recipeStore := db.NewStore(s.db)
	recipeHandler := recipe.NewHandler(recipeStore)
	recipeHandler.RegisterRoutes(router)

	log.Println("Server running on: ", s.addr)

	return http.ListenAndServe(s.addr, router)
}
