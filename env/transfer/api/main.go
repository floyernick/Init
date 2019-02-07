package api

import (
	"Init/env/config"
	"Init/usecases"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"golang.org/x/crypto/sha3"
)

type API struct {
	controller usecases.Controller
	checkHash  bool
	hashSalt   string
}

func NewAPIHandler(config config.HandlerConfig, controller usecases.Controller) *http.ServeMux {
	api := API{controller, config.CheckHash, config.HashSalt}
	mux := http.NewServeMux()
	mux.HandleFunc("/notes.create", api.NotesCreate)
	mux.HandleFunc("/notes.update", api.NotesUpdate)
	mux.HandleFunc("/notes.delete", api.NotesDelete)
	mux.HandleFunc("/notes.get", api.NotesGet)
	return mux
}

func ProcessRequest(request *http.Request, v interface{}, checkHash bool, hashSalt string) error {

	requestBody, err := ioutil.ReadAll(request.Body)

	if err != nil {
		return errors.New("invalid request data")
	}

	defer request.Body.Close()

	err = json.Unmarshal(requestBody, v)

	if err != nil {
		return errors.New("invalid request data")
	}

	if checkHash {

		requestHash := request.Header.Get("X-Hash")

		hashSequence := make([]byte, len(requestBody), len(requestBody)*2)
		copy(hashSequence, requestBody)
		hashSequence = append(hashSequence, byte(len(hashSequence)%255))
		hashSequence = append(hashSequence, []byte(hashSalt)...)
		hashValue := sha3.Sum256(hashSequence)

		if requestHash != hex.EncodeToString(hashValue[:]) {
			return errors.New("invalid request hash")
		}

	}

	return nil

}

type Response struct {
	Status bool        `json:"status"`
	Result interface{} `json:"result,omitempty"`
	Error  interface{} `json:"error,omitempty"`
}

func returnResponse(response http.ResponseWriter, data []byte) {

	response.Header().Set("Content-Type", "application/json")
	response.Header().Set("Access-Control-Allow-Origin", "*")
	response.Header().Set("Access-Control-Allow-Headers", "X-Hash, Content-Type")
	response.Write(data)
}

func ReturnSuccessResponse(response http.ResponseWriter, v interface{}) {
	responseBody := Response{
		Status: true,
		Result: v,
	}
	data, _ := json.Marshal(responseBody)
	returnResponse(response, data)
}

func ReturnErrorResponse(response http.ResponseWriter, v interface{}) {
	responseBody := Response{
		Status: false,
		Error:  v,
	}
	data, _ := json.Marshal(responseBody)
	returnResponse(response, data)
}
