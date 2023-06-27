package handlers

import (
	"net/http"

	"github.com/anjotadena/centauri"
)

type Handlers struct {
	App *centauri.Centauri
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	err := h.App.Render.Page(w, r, "home", nil, nil)

	if err != nil {
		h.App.ErrorLog.Println("Error rendering: ", err)
	}
}
