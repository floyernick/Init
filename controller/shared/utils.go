package presenter

import (
	"Init/app/errors"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseRequestBody(request *http.Request, v interface{}) error {

	requestBody, err := ioutil.ReadAll(request.Body)

	if err != nil {
		return errors.BadRequest
	}

	defer request.Body.Close()

	err = json.Unmarshal(requestBody, v)

	if err != nil {
		return errors.BadRequest
	}

	return nil

}

type responseError struct {
	Error string `json:"error"`
}

func RespondWithSuccess(response http.ResponseWriter, v interface{}) {
	data, _ := json.Marshal(v)
	respond(response, data)
}

func RespondWithError(response http.ResponseWriter, err error) {
	error := responseError{Error: err.Error()}
	data, _ := json.Marshal(error)
	respond(response, data)
}

func respond(response http.ResponseWriter, data []byte) {

	response.Header().Set("Content-Type", "application/json")
	response.Header().Set("Access-Control-Allow-Origin", "*")
	response.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	response.Write(data)
}
