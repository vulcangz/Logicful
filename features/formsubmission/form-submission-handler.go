package formsubmission

import (
	"api/handler"
	_ "encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/logicful/models"
	"github.com/logicful/service/httpextensions"
	"net/http"
)

func AddHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var submission = models.Submission{}
	if !httpextensions.ReadJson(&submission, w, r) {
		return
	}
	submission.Meta.IpAddress = r.RemoteAddr
	submission.FormId = ps.ByName("formId")
	err := Add(submission)
	if err != nil {
		httpextensions.WriteError(w, err)
		return
	}
	httpextensions.WriteNoContent(w)
}

func GetReadHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	read, err := GetRead(ps.ByName("formId"), handler.User(r))
	if err != nil {
		httpextensions.WriteError(w, err)
		return
	}
	httpextensions.WriteJson(w, read)
}

func MarkReadHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var body map[string]bool
	if !httpextensions.ReadJson(&body, w, r) {
		return
	}
	err := MarkRead(body, ps.ByName("formId"), handler.User(r))
	if err != nil {
		httpextensions.WriteError(w, err)
		return
	}
	httpextensions.WriteNoContent(w)
}

func ListHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	url, err := List(ps.ByName("formId"), handler.User(r))
	if err != nil {
		httpextensions.WriteError(w, err)
		return
	}
	httpextensions.WriteJson(w, httpextensions.SuccessResult{
		Message: url,
	})
}

func GenerateFileUrlHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var file = models.File{}
	if !httpextensions.ReadJson(&file, w, r) {
		return
	}
	url, err := GetSubmissionFile(file, handler.User(r))
	if err != nil {
		httpextensions.WriteError(w, err)
		return
	}
	httpextensions.WriteJson(w, httpextensions.SuccessResult{
		Message: url,
	})
}

func DeleteHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var ids []string
	if !httpextensions.ReadJson(&ids, w, r) {
		return
	}
	err := Delete(ids, ps.ByName("formId"), handler.User(r))
	if err != nil {
		httpextensions.WriteError(w, err)
		return
	}
	httpextensions.WriteNoContent(w)
}
