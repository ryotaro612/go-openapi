package main

import (
	"github.com/ryotaro612/go-openapi/internal/handler"
	"net/http"
	"time"
)

func main() {
	router := handler.NewRouter()
	s := http.Server{
		Addr:         ":8080",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      router,
	}
	err := s.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}
}
