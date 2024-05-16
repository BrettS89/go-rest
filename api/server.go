package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	listenAddr string
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
	}
}

func registerRoutes() *mux.Router {
	router := mux.NewRouter()

	userRoutes := NewUserRoutes()
	userRoutes.RegisterRoutes(router)

	router.HandleFunc("/", helloWorldHandler)

	return router
}

func (s *Server) Start() error {
	log.Println("Server listening on port 4000!")

	router := registerRoutes()

	return http.ListenAndServe(s.listenAddr, router)
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	response := map[string]string{
		"message": "hello world!",
	}

	json.NewEncoder(w).Encode(response)
}
