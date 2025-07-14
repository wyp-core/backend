package transport

import (
	"context"
	"encoding/json"

	"net/http"

	"github.com/Abhyuday04/wyp/constants"
	responsehandler "github.com/Abhyuday04/wyp/handlers/responseHandler"
	"github.com/Abhyuday04/wyp/layers/models"
	"github.com/Abhyuday04/wyp/layers/services"
	"github.com/rs/zerolog/log"
)

type Transport struct {
	Service services.IService
}

func New(service services.IService) *Transport {
	return &Transport{
		Service: service,
	}
}

func (t *Transport) AddUserCont(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()
		var reqObj models.AddUserFetchParam
		err := decoder.Decode(&reqObj)
		if err != nil {
			log.Error().Err(err)
			responsehandler.GenericErrRes(w, http.StatusBadRequest, constants.ErrorInvalidRequestBody)
			return
		}
		respParams, err := t.Service.AddUser(&reqObj)
		if err != nil {
			log.Error().Err(err)
			responsehandler.GenericErrRes(w, http.StatusInternalServerError, err.Error())
			return
		}

		ctx := context.WithValue(r.Context(), "resData", respParams)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (t *Transport) AddJobCont(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()
		var job models.Job
		err := decoder.Decode(&job)
		// TODO: Handle error properly
		if err != nil {
			log.Error().Err(err)
			responsehandler.GenericErrRes(w, http.StatusBadRequest, constants.ErrorInvalidRequestBody)
			return
		}
		respParams, err := t.Service.AddJob(&job)
		if err != nil {
			log.Error().Err(err)
			responsehandler.GenericErrRes(w, http.StatusInternalServerError, err.Error())
			return
		}

		ctx := context.WithValue(r.Context(), "resData", respParams)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (t *Transport) GetJobsCont(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()
		var fetchParams models.GetJobsFetchParam
		err := decoder.Decode(&fetchParams)
		// TODO: Handle error properly
		if err != nil {
			log.Error().Err(err)
			responsehandler.GenericErrRes(w, http.StatusBadRequest, constants.ErrorInvalidRequestBody)
			return
		}
		respParams, err := t.Service.GetJobs(&fetchParams)
		if err != nil {
			log.Error().Err(err)
			responsehandler.GenericErrRes(w, http.StatusInternalServerError, err.Error())
			return
		}

		ctx := context.WithValue(r.Context(), "resData", respParams)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (t *Transport) SendOtpCont(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()
		var reqObj models.SendOtpParam
		err := decoder.Decode(&reqObj)
		if err != nil {
			log.Error().Err(err).Msg("Failed to decode request body")
			responsehandler.GenericErrRes(w, http.StatusBadRequest, constants.ErrorInvalidRequestBody)
			return
		}
		if reqObj.CountryCode == "" || reqObj.Phone == "" {
			log.Error().Msg("CountryCode or Phone is empty")
			responsehandler.GenericErrRes(w, http.StatusBadRequest, "CountryCode or Phone is empty")
			return
		}
		err = t.Service.SendOtp(r.Context(), &reqObj)
		if err != nil {
			log.Error().Err(err)
			responsehandler.GenericErrRes(w, http.StatusInternalServerError, err.Error())
			return
		} else {
			log.Info().Msg("OTP sent successfully")
		}

		ctx := context.WithValue(r.Context(), "resData", "Successfully sent OTP")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (t *Transport) VerifyOtpCont(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()
		var reqObj models.VerifyOtpParam
		err := decoder.Decode(&reqObj)
		if err != nil {
			log.Error().Err(err)
			responsehandler.GenericErrRes(w, http.StatusBadRequest, constants.ErrorInvalidRequestBody)
			return
		}
		userID, err := t.Service.VerifyOtp(r.Context(), &reqObj)
		if err != nil {
			log.Error().Err(err)
			responsehandler.GenericErrRes(w, http.StatusInternalServerError, err.Error())
			return
		}

		ctx := context.WithValue(r.Context(), "resData", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
