package api

import (
	"github.com/go-chi/chi/v5"
)

func initRouters(s *APIServer) *chi.Mux {
	r := chi.NewRouter()
	r.Route("/api/v1/account", func(r chi.Router) {
		r.Get("/{id}", makeHttpHandleFunc(s.handleGetAccountById))
		r.Get("/all", makeHttpHandleFunc(s.handleGetAccounts))
		r.Post("/", makeHttpHandleFunc(s.handleCreateAccount))
		r.Post("/transfer", makeHttpHandleFunc(s.handleTransfer))
		r.Put("/{id}", makeHttpHandleFunc(s.handleUpdateAccount))
		r.Delete("/{id}", makeHttpHandleFunc(s.handleDeleteAccount))
	})
	return r
}
