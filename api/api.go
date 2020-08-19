package api

import (
	"github.com/railgun-project/railgun/assets"
	"net/http"
	"time"
)

func Run() {
	_ = assets.Load("./assets/")
	srv := &http.Server{
		Handler: routes(),
		Addr:    ":8001",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	_ = srv.ListenAndServe()
}

func RunTLS(certFile, keyFile string) {
	srv := &http.Server{
		Handler: routes(),
		Addr:    ":8001",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	_ = srv.ListenAndServeTLS(certFile, keyFile)
}
