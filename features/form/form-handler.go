package form

import (
	"api/handler"
	_ "encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/logicful/models"
	"github.com/logicful/service/httpextensions"
	"net/http"
)

func SetHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var form = models.Form{}
	if !httpextensions.ReadJson(&form, w, r) {
		return
	}
	form.Id = ps.ByName("formId")
	form, err := Set(form, handler.User(r))
	if err != nil {
		httpextensions.WriteError(w, err)
		return
	}
	httpextensions.WriteJson(w, form)
}

func ListHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	forms, err := List(httpextensions.Query("folderId", r))
	if err != nil {
		httpextensions.WriteError(w, err)
		return
	}
	httpextensions.WriteJson(w, forms)
}

func GetHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	form, err := Get(ps.ByName("formId"))
	if err != nil {
		httpextensions.WriteError(w, err)
		return
	}
	httpextensions.WriteJson(w, form)
}
