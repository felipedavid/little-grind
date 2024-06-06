package main

import (
	"log/slog"
	"net/http"
	"waifus/handlers"
)

func main() {
	routes := handlers.SetupRoutes()

	s := &http.Server{
		Addr:    ":3000",
		Handler: routes,
	}

	slog.Info("Server starting", "addr", s.Addr)
	err := s.ListenAndServe()
	slog.Error("Server critical error", "err", err)
	panic(err)
}
