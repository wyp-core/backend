
package v1

import (
	responsehandler "github.com/Abhyuday04/wyp/handlers/responseHandler"
	"github.com/Abhyuday04/wyp/internal/app"
	"github.com/go-chi/chi/v5"
)

func JobRouter(r chi.Router) {
	r.With(
		app.Srv.Transport.AddJobCont,
	).Post("/addJob", responsehandler.GenericRes)
}