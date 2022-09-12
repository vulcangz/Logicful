package contentblock

import (
	_ "encoding/json"
	"logicful/models"
	"logicful/service/httpextensions"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func SetHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var block = models.ContentBlock{}
	if !httpextensions.ReadJson(&block, w, r) {
		return
	}
	block.Id = ps.ByName("id")
	err := Set(block)
	if err != nil {
		httpextensions.WriteError(w, err)
		return
	}
	httpextensions.WriteNoContent(w)
}

func ListHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sets, err := List()
	if err != nil {
		httpextensions.WriteError(w, err)
		return
	}
	httpextensions.WriteJson(w, sets)
}
