package api

import (
	"encoding/json"
	"net/http"

	"github.com/BrettS89/models"
	"github.com/gorilla/mux"
)

type UserRoutes struct {
	Path string
}

func NewUserRoutes() *UserRoutes {
	return &UserRoutes{Path: "/user"}
}

func (u *UserRoutes) RegisterRoutes(r *mux.Router) {
	r.HandleFunc(u.Path, WrapperHandler(endpointsWithoutParam))
	r.HandleFunc(u.Path+"/{id}", WrapperHandler(endpointsWithIDParam))
}

func endpointsWithIDParam(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return handleGetUser(w, r)
	}

	if r.Method == "PATCH" {
		return handlePatchUser(w, r)
	}

	return nil
}

func endpointsWithoutParam(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST" {
		return handleCreateUser(w, r)
	}

	return nil
}

func handleCreateUser(w http.ResponseWriter, r *http.Request) error {
	u := models.NewUser()

	user, err := u.Create(r.Body)

	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(user)

	return nil
}

func handleGetUser(w http.ResponseWriter, r *http.Request) error {
	userId := mux.Vars(r)["id"]

	u := models.NewUser()

	user, err := u.Get(userId)

	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(user)

	return nil
}

func handlePatchUser(w http.ResponseWriter, r *http.Request) error {
	userId := mux.Vars(r)["id"]

	u := models.NewUser()

	user, err := u.Patch(userId, r.Body)

	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(user)

	return nil
}
