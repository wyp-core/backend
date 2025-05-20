package transport

import "net/http"

type ITransport interface {
	AddUserCont(next http.Handler) http.Handler
}