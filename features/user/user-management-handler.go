package user

import (
	"logicful/models"
	"logicful/service/httpextensions"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var user = models.User{}
	if !httpextensions.ReadJson(&user, w, r) {
		return
	}
	token, err := Register(user, false)
	if err != nil {
		httpextensions.WriteError(w, err)
		return
	}
	httpextensions.WriteJson(w, token)
}

func LoginHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var user = models.User{}
	if !httpextensions.ReadJson(&user, w, r) {
		return
	}
	token, err := Login(user)
	if err != nil {
		httpextensions.WriteError(w, err)
		return
	}
	httpextensions.WriteJson(w, token)
}

func LoginFromGoogleHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var request = models.TokenRequest{}
	if !httpextensions.ReadJson(&request, w, r) {
		return
	}
	token, err := LoginFromGoogleToken(request.Token)
	if err != nil {
		httpextensions.WriteError(w, err)
		return
	}
	httpextensions.WriteJson(w, token)
}

func RefreshHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var token = r.Header.Get("Authorization")
	token = strings.Replace(token, "Bearer ", "", 1)
	token = strings.TrimSpace(token)
	refreshed := RefreshToken(token)
	httpextensions.WriteJson(w, refreshed)
}
