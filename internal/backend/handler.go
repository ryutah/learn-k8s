package backend

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

func Handler() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.StripSlashes)

	h := new(handler)
	r.Post("/", h.handle)

	return r
}

type payload struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type errorResponse struct {
	Message string
}

type handler struct{}

func (*handler) handle(w http.ResponseWriter, r *http.Request) {
	var p payload
	if err := render.Decode(r, &p); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, errorResponse{
			Message: err.Error(),
		})
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, p)
}
