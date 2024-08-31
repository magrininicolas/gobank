package api

import (
	"github.com/go-chi/chi/v5"
)

func initRouters(s *APIServer) *chi.Mux {
	r := chi.NewRouter()
	r.Route("/api/v1/account", func(r chi.Router) {
		r.Get("/{id}", makeHttpHandleFunc(s.handleGetAccount))
		r.Post("/transfer", makeHttpHandleFunc(s.handleTransfer))
		r.Post("/", makeHttpHandleFunc(s.handleCreateAccount))
		r.Delete("/{id}", makeHttpHandleFunc(s.handleDeleteAccount))
		r.Put("/{id}", makeHttpHandleFunc(s.handleUpdateAccount))
	})
	return r
}
