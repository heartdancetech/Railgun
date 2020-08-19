package api

import (
	"github.com/gorilla/mux"
	"github.com/railgun-project/railgun/assets"
	"net/http"
)

func routes() *mux.Router {
	r := mux.NewRouter()
	r.PathPrefix("/api").Path("/get_proxy").HandlerFunc(handle).Methods("GET")
	r.PathPrefix("/api").Path("/put_proxy").HandlerFunc(handle).Methods("PUT")

	r.Handle("/favicon.ico", http.FileServer(assets.FileSystem)).Methods("GET")
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(assets.FileSystem))).Methods("GET")
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/static", http.StatusMovedPermanently)
	})
	return r
}
