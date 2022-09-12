package folder

import (
	_ "encoding/json"
	"logicful/handler"
	"logicful/models"
	"logicful/service/httpextensions"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func SetHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var folder = models.Folder{}
	if !httpextensions.ReadJson(&folder, w, r) {
		return
	}
	folder.Id = ps.ByName("folderId")
	folder.TeamId = handler.User(r).TeamId
	folder, err := Set(folder, handler.User(r))
	if err != nil {
		httpextensions.WriteError(w, err)
		return
	}
	httpextensions.WriteJson(w, folder)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := Delete(ps.ByName("folderId"), handler.User(r))
	if err != nil {
		httpextensions.WriteError(w, err)
		return
	}
	httpextensions.WriteNoContent(w)
}

func ListHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	folders, err := List(handler.User(r).TeamId)
	if err != nil {
		httpextensions.WriteError(w, err)
		return
	}
	httpextensions.WriteJson(w, folders)
}
