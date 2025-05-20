package transport

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Abhyuday04/wyp/layers/models"
	"github.com/Abhyuday04/wyp/layers/services"
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
			log.Println(err)
			panic(err)
		}
		respParams ,err := t.Service.AddUser(&reqObj)	

		ctx := context.WithValue(r.Context(), "resData", respParams)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
