package handlers

import (
	"net/http"
	"waifus/views/home"
)

func RenderHomePage(w http.ResponseWriter, r *http.Request) error {
	return home.HomePage().Render(r.Context(), w)
}
