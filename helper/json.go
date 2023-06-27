package helper

import (
	"encoding/json"
	"net/http"
)

// create a function that decodes the request and result body
func ReadRequestBody(r *http.Request, result interface{}) {
	// decode the request body
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(result)
	PanicIfError(err) // if the function/file exist within the same package, no need of the package name

}

func WriteResponseBody(w http.ResponseWriter, response interface{}) {
	// set the response header type
	w.Header().Add("Content-Type", "Application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)
	PanicIfError(err)
}
