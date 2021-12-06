package api

import (
	"github.com/heart-dance-x/railgun/assets"
	_ "github.com/heart-dance-x/railgun/assets/statik"
	"net/http"
	"time"
)

func init() {
	_ = assets.Load("")
}

func Run() {
	srv := &http.Server{
		Handler:      routes(),
		Addr:         ":8001",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	_ = srv.ListenAndServe()
}

func RunTLS(certFile, keyFile string) {
	srv := &http.Server{
		Handler:      routes(),
		Addr:         ":8001",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	_ = srv.ListenAndServeTLS(certFile, keyFile)
}
