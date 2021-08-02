package optionset

import (
	"api/handler"
	_ "encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/logicful/models"
	"github.com/logicful/service/httpextensions"
	"net/http"
)

func SetHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var set = models.OptionSet{}
	set.TeamId = handler.User(r).TeamId
	if !httpextensions.ReadJson(&set, w, r) {
		return
	}
	set.Id = ps.ByName("id")
	err := Set(set, handler.User(r))
	if err != nil {
		httpextensions.WriteError(w, err)
		return
	}
	httpextensions.WriteNoContent(w)
}

func ListHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sets, err := List(handler.User(r).TeamId)
	if err != nil {
		httpextensions.WriteError(w, err)
		return
	}
	httpextensions.WriteJson(w, sets)
}
