package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"time"

	"github.com/ExonegeS/REST-API-001/internal/domain"
	"github.com/gorilla/mux"
)

type ApiServer struct {
	svc domain.Service
	srv *http.Server
}

func NewApiServer(svc domain.Service) *ApiServer {
	return &ApiServer{
		svc: svc,
	}
}

func (s *ApiServer) Start(listenAddr int) error {
	router := mux.NewRouter()
	s.srv = &http.Server{
		Addr:    fmt.Sprintf(":%v", listenAddr),
		Handler: router,
	}
	router.HandleFunc("/users", s.getUsersHandler).Methods("GET")
	router.HandleFunc("/users/{id}", s.getUserHandler).Methods("GET")
	router.HandleFunc("/users", s.createUserHandler).Methods("POST")
	router.HandleFunc("/users/{id}", s.updateUserHandler).Methods("PUT")

	slog.Info(fmt.Sprintf("SERVER STARTED AT ADDRESS %v", listenAddr))
	return s.srv.ListenAndServe()
}

func (s *ApiServer) Stop(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := s.srv.Shutdown(ctx); err != nil {
		log.Println("HTTP server shutdown failed:", err)
	} else {
		log.Println("Server stopped gracefully")
	}
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func (s *ApiServer) getUsersHandler(w http.ResponseWriter, r *http.Request) {
	limit, offset, orderBy, query, err := parseQueryParams(r)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		return
	}

	input := domain.GetUsersInput{
		Limit:   limit,
		Offset:  offset,
		OrderBy: orderBy,
		Query:   query,
	}

	response, err := s.svc.GetUsersMany(context.Background(), input)
	if err != nil {
		writeJSON(w, http.StatusConflict, map[string]any{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, response)
}

func (s *ApiServer) getUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": "ID is required"})
		return
	}

	user, err := s.svc.GetUsersOne(context.Background(), domain.GetUserInput{
		ID: &id,
	})
	if err != nil {
		writeJSON(w, http.StatusNotFound, map[string]any{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, user)
}

func (s *ApiServer) createUserHandler(w http.ResponseWriter, r *http.Request) {
	var input domain.CreateUserInput

	if r.Body == nil || r.ContentLength == 0 {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": "Request body is required"})
		return
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		if err == io.EOF {
			writeJSON(w, http.StatusBadRequest, map[string]any{"error": "Empty JSON body"})
		} else {
			writeJSON(w, http.StatusBadRequest, map[string]any{"error": "Invalid JSON body"})
		}
		slog.Error(err.Error(), "input", input)
		return
	}

	response, err := s.svc.CreateUser(context.Background(), input)
	if err != nil {
		writeJSON(w, http.StatusConflict, map[string]any{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, response)
}

func (s *ApiServer) updateUserHandler(w http.ResponseWriter, r *http.Request) {
	var input domain.UpdateUserInput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": "Invalid JSON body"})
		return
	}

	userID := mux.Vars(r)["id"]
	input.ID = userID

	updatedUser, err := s.svc.UpdateUser(context.Background(), input)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, updatedUser)
}
