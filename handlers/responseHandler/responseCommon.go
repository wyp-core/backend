package responsehandler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"syscall"
)

// add recovery
func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err = w.Write(response)
	// This error check is to handle the write broken pipe error.
	// This happens if the client closed the connection,
	// but the server still tries to send more data to the client
	// and receives an RST packet indicating that the connection was closed.
	// This error can be safely ignored.
	if errors.Is(err, syscall.EPIPE) {
		log.Println(err)
		return
	}
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

func GenericRes(w http.ResponseWriter, r *http.Request) {

	resData := r.Context().Value("resData")

	var payload = map[string]interface{}{
		"status": true,
		"error":  "",
		"data":   resData,
	}
	respondwithJSON(w, http.StatusOK, payload)

}

func GenericErrRes(w http.ResponseWriter, code int, payload interface{}) {

	var errJson = map[string]interface{}{
		"status": false,
		"error":  payload,
		"data":   nil,
	}
	respondwithJSON(w, http.StatusInternalServerError, errJson)
}