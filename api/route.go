package api

import (
	"github.com/gorilla/mux"
)

func routes() *mux.Router {
	r := mux.NewRouter()
	r.PathPrefix("/api").Path("/get_proxy").HandlerFunc(handle).Methods("GET")
	r.PathPrefix("/api").Path("/put_proxy").HandlerFunc(handle).Methods("PUT")
	return r
}
