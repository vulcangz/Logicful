package s3store

import (
	"errors"
	"io/ioutil"
	"logicful/service/httpextensions"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func SetJsonHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		httpextensions.WriteError(w, err)
		return
	}
	url, err := SetJson(httpextensions.Query("id", r), bytes)
	if err != nil {
		httpextensions.WriteError(w, err)
		return
	}
	httpextensions.WriteJson(w, httpextensions.SuccessResult{
		Message: url,
	})
}

func GenerateUrlHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	length := httpextensions.Query("length", r)
	if length == "" {
		httpextensions.WriteError(w, errors.New("length must be supplied"))
		return
	}
	parsed, err := strconv.ParseInt(length, 10, 64)
	if err != nil {
		httpextensions.WriteError(w, err)
		return
	}
	if parsed >= 5e+7 {
		httpextensions.WriteError(w, errors.New("file may not be greater than 50 mb"))
		return
	}
	url, key, err := GenerateStoreUrl(parsed)
	if err != nil {
		httpextensions.WriteError(w, err)
		return
	}
	type UrlResult struct {
		Key string `json:"key"`
		Url string `json:"url"`
	}
	httpextensions.WriteJson(w, UrlResult{
		Key: key,
		Url: url,
	})
}
