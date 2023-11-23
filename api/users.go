package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"webservermod/model"

	"github.com/go-chi/chi/v5"
)

func (s *server) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	// Implement the user creation logic
	body, err := io.ReadAll(r.Body)
	if err != nil {
		sendApiError(w, "Unable to read the body t", http.StatusBadRequest)
	}
	user := model.User{}
	if err := json.Unmarshal(body, &user); err != nil {
		sendApiError(w, "Bad request", http.StatusBadRequest)
		return
	}

	if err := s.userDB.CreateUser(user); err != nil {
		sendApiError(w, "Internal server error ", http.StatusInternalServerError)
		return
	}

	sendApiError(w, "User Created successfully", http.StatusCreated)

}

func (s *server) HandleGetByEmail(w http.ResponseWriter, r *http.Request) {

	// Implement the user creation logic
	email := chi.URLParam(r, "email")
	if email == "" {
		sendApiError(w, "Email id not exists ", http.StatusBadRequest)
	}
	user, err := s.userDB.GetUserByEmail(email)
	if err != nil {
		sendApiError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(user)
	if err != nil {
		sendApiError(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func sendApiError(w http.ResponseWriter, msg string, code int) {
	apiError := model.ApiError{Msg: msg, Code: code}
	bytes, err := json.Marshal(&apiError)
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(code)
	w.Write(bytes)
}
