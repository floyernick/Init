package presenter

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func handleRequest(handler func(r *http.Request) (interface{}, error)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := handler(r)
		if err != nil {
			respondWithError(w, err.Error())
		} else {
			respondWithSuccess(w, result)
		}
	}
}

func parseRequestBody(request *http.Request, v interface{}) error {

	requestBody, err := ioutil.ReadAll(request.Body)

	if err != nil {
		return errors.New("invalid request data")
	}

	defer request.Body.Close()

	err = json.Unmarshal(requestBody, v)

	if err != nil {
		return errors.New("invalid request data")
	}

	return nil

}

type responseBody struct {
	Status bool        `json:"status"`
	Result interface{} `json:"result,omitempty"`
	Error  interface{} `json:"error,omitempty"`
}

func respondWithSuccess(response http.ResponseWriter, v interface{}) {
	responseBody := responseBody{
		Status: true,
		Result: v,
	}
	data, _ := json.Marshal(responseBody)
	respond(response, data)
}

func respondWithError(response http.ResponseWriter, v interface{}) {
	responseBody := responseBody{
		Status: false,
		Error:  v,
	}
	data, _ := json.Marshal(responseBody)
	respond(response, data)
}

func respond(response http.ResponseWriter, data []byte) {

	response.Header().Set("Content-Type", "application/json")
	response.Header().Set("Access-Control-Allow-Origin", "*")
	response.Header().Set("Access-Control-Allow-Headers", "X-Hash, Content-Type")
	response.Write(data)
}
