
package v1

import (
	responsehandler "github.com/Abhyuday04/wyp/handlers/responseHandler"
	"github.com/Abhyuday04/wyp/internal/app"
	"github.com/go-chi/chi/v5"
)

func JobRouter(r chi.Router) {
	r.With(
		app.Srv.Transport.AddJobCont,
	).Post("/", responsehandler.GenericRes)
	r.With(
		app.Srv.Transport.GetJobsCont,
	).Post("/all", responsehandler.GenericRes)
}