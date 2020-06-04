package presenter

import (
	"Init/app/errors"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func handleRequest(handler func(r *http.Request) (interface{}, error)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := handler(r)
		if err != nil {
			respondWithError(w, err)
		} else {
			respondWithSuccess(w, result)
		}
	}
}

func parseRequestBody(request *http.Request, v interface{}) error {

	requestBody, err := ioutil.ReadAll(request.Body)

	if err != nil {
		return errors.BadRequest{}
	}

	defer request.Body.Close()

	err = json.Unmarshal(requestBody, v)

	if err != nil {
		return errors.BadRequest{}
	}

	return nil

}

type ResponseError struct {
	Error string `json:"error"`
}

func respondWithSuccess(response http.ResponseWriter, v interface{}) {
	data, _ := json.Marshal(v)
	respond(response, data)
}

func respondWithError(response http.ResponseWriter, err error) {
	error := ResponseError{Error: err.Error()}
	data, _ := json.Marshal(error)
	respond(response, data)
}

func respond(response http.ResponseWriter, data []byte) {

	response.Header().Set("Content-Type", "application/json")
	response.Header().Set("Access-Control-Allow-Origin", "*")
	response.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	response.Write(data)
}
