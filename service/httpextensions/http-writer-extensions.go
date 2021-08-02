package httpextensions

import (
	"encoding/json"
	"net/http"
)

type ErrorResult struct {
	Message string `json:"message"`
}

type SuccessResult struct {
	Message string `json:"message"`
}

const MaxAgeOneWeek = 604800
const MaxAgeOneMinute = 60

func WriteJson(writer http.ResponseWriter, data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		WriteError(writer, err)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writeCorsHeaders(writer)
	_, err = writer.Write(bytes)
	if err != nil {
		panic(err)
	}
}

func WriteNoContent(writer http.ResponseWriter) {
	writeCorsHeaders(writer)
	writer.WriteHeader(204)
}

func WriteError(writer http.ResponseWriter, error error) {
	writeCorsHeaders(writer)
	writer.WriteHeader(400)
	WriteJson(writer, ErrorResult{
		Message: error.Error(),
	})
}

func writeCorsHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Request-Method", "*")
	writer.Header().Set("Access-Control-Request-Headers", "*")
}
