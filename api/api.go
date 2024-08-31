package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type apiFunc func(http.ResponseWriter, *http.Request) error

type JSONResponse map[string]any

type APIError struct {
	Error string
}

type APIServer struct {
	listenAddr string
}

func makeHttpHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
		}
	}
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func NewApiServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func (s *APIServer) Run() {
	r := initRouters(s)

	log.Printf("JSON API server running on port: %s\n", s.listenAddr)

	http.ListenAndServe(s.listenAddr, r)
}
