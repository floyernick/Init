package api

import (
	"Init/usecases"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type API struct {
	controller usecases.Controller
}

func NewAPIHandler(controller usecases.Controller) *http.ServeMux {
	api := &API{controller}
	mux := http.NewServeMux()
	mux.HandleFunc("/users.get", api.UsersGet)
	return mux
}

func ProcessRequest(request *http.Request, v interface{}) error {

	requestBody, err := ioutil.ReadAll(request.Body)

	if err != nil {
		return err
	}

	defer request.Body.Close()

	err = json.Unmarshal(requestBody, v)

	if err != nil {
		return err
	}

	return nil

}

func returnResponse(response http.ResponseWriter, data []byte) {

	response.Header().Set("Content-Type", "application/json")
	response.Header().Set("Access-Control-Allow-Origin", "*")
	response.Header().Set("Access-Control-Allow-Headers", "X-Hash, Content-Type")
	response.Write(data)
}

func ReturnSuccessResponse(response http.ResponseWriter, v interface{}) {
	responseBody := map[string]interface{}{
		"status": true,
		"result": v,
	}
	data, _ := json.Marshal(responseBody)
	returnResponse(response, data)
}

func ReturnErrorResponse(response http.ResponseWriter, v interface{}) {
	responseBody := map[string]interface{}{
		"status": false,
		"error":  v,
	}
	data, _ := json.Marshal(responseBody)
	returnResponse(response, data)
}
