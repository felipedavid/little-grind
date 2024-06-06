package handlers

import (
	"log/slog"
	"net/http"
	"waifus/assets"
)

func SetupRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.FS(assets.AssetsFS))))

	defineRoute(mux, "/", RenderHomePage)

	return mux
}

type customHandler func(w http.ResponseWriter, r *http.Request) error

func defineRoute(mux *http.ServeMux, path string, h customHandler) {
	mux.HandleFunc(path, stdlibHandler(h))
}

func stdlibHandler(h customHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := h(w, r)
		if err != nil {
			slog.Error("Handler returned an error", "err", err)
		}
	}
}
