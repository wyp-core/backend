package v1

import "github.com/go-chi/chi/v5"

func V1Routes(r chi.Router) {
	r.Route("/user", UserRouter)
}
