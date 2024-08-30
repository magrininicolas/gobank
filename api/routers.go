package api

import "github.com/gorilla/mux"

func initRouters(s *APIServer) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/account", makeHttpHandleFunc(s.handleAccount))
	router.HandleFunc("/account/{id}", makeHttpHandleFunc(s.handleGetAccount))
	router.HandleFunc("/api/v1/account/transfer", makeHttpHandleFunc(s.handleTransfer))

	return router
}
