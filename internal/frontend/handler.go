package frontend

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Handler() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.StripSlashes)

	h := new(handler)
	r.Get("/", h.index)

	return r
}

type handler struct{}

func (*handler) index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/frontend/index.html")
}
