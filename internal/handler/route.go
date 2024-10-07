package handler

import (
	"github.com/gorilla/mux"
	"net/http"
)

// NewRouter creates a new router.
func NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.Handle("/v1/health", newHealthHandler()).Methods(http.MethodGet)
	router.Handle("/scim/v2/Users", newUserCreationHandler()).Methods(http.MethodPost)
	router.Handle("/scim/v2/Users/{id}", newUserUpdateHandler()).Methods(http.MethodPut)
	return router
}
