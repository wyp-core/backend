package v1

import (
	responsehandler "github.com/Abhyuday04/wyp/handlers/responseHandler"
	"github.com/Abhyuday04/wyp/internal/app"
	"github.com/go-chi/chi/v5"
)

func OtpRouter(r chi.Router) {
	r.With(
		app.Srv.Transport.SendOtpCont,
	).Post("/send", responsehandler.GenericRes)
	r.With(
		app.Srv.Transport.VerifyOtpCont,
	).Post("/verify", responsehandler.GenericRes)
}