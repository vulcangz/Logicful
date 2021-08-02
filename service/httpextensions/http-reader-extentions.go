package httpextensions

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Query(key string, r *http.Request) string {
	return r.URL.Query().Get(key)
}

func ReadJson(value interface{}, w http.ResponseWriter, r *http.Request) bool {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		WriteError(w, err)
		return false
	}

	err = json.Unmarshal(b, &value)
	if err != nil {
		WriteError(w, err)
		return false
	}
	return true
}
