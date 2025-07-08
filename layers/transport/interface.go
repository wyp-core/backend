package transport

import "net/http"

type ITransport interface {
	AddUserCont(next http.Handler) http.Handler
	AddJobCont(next http.Handler) http.Handler
	GetJobsCont(next http.Handler) http.Handler
	SendOtpCont(next http.Handler) http.Handler
	VerifyOtpCont(next http.Handler) http.Handler
}