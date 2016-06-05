package res

import (
	"encoding/json"
	"net/http"
)

type responseError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type data struct {
	Item interface{} `json:"item"`
}

// This properties are interface{} because
// if any of it is missing it should be null,
// not the value which is given to the struct when
// it is initialized empty
type response struct {
	Data  interface{} `json:"data"`
	Error interface{} `json:"error"`
}

type status struct {
	StatusCode int
	StatusText string
}

func sendOutJs(w http.ResponseWriter, status status, js []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status.StatusCode)
	w.Write(js)
}

// Finalize is a function which can send back
// structured response to the client
func Finalize(w http.ResponseWriter, object interface{}) {
	finalData := data{
		Item: object,
	}
	final := response{
		Data: finalData,
	}
	js, err := json.Marshal(final)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	responseStatus := status{
		StatusCode: http.StatusOK,
		StatusText: http.StatusText(http.StatusOK),
	}
	sendOutJs(w, responseStatus, js)
}

// FinalizeError is used to send out structured errors
// to the client
func FinalizeError(w http.ResponseWriter, object error, statusCode int) {
	finalError := responseError{
		Message: object.Error(),
		Status:  statusCode,
	}
	final := response{
		Error: finalError,
	}

	js, err := json.Marshal(final)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	responseStatus := status{
		StatusCode: statusCode,
		StatusText: http.StatusText(statusCode),
	}
	sendOutJs(w, responseStatus, js)
}
