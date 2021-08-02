package team

import (
	"api/features/user"
	"api/handler"
	_ "encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/logicful/models"
	"github.com/logicful/service/httpextensions"
	"net/http"
)

func TeamHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	token, err := ByUser(handler.User(r))
	if err != nil {
		httpextensions.WriteError(w, err)
		return
	}
	httpextensions.WriteJson(w, token)
}

func AcceptInviteHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := AcceptInvite(handler.User(r), httpextensions.Query("team", r))
	if err != nil {
		httpextensions.WriteError(w, err)
		return
	}
	httpextensions.WriteNoContent(w)
}

func MembersHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	token, err := user.ByTeam(handler.User(r))
	if err != nil {
		httpextensions.WriteError(w, err)
		return
	}
	httpextensions.WriteJson(w, token)
}

func RemoveMembersHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var request = models.RemoveTeammatesRequest{}
	if !httpextensions.ReadJson(&request, w, r) {
		return
	}
	err := RemoveMembers(handler.User(r), request.UserIds)
	if err != nil {
		httpextensions.WriteError(w, err)
		return
	}
	httpextensions.WriteNoContent(w)
}

func InviteMembersHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var request []string
	if !httpextensions.ReadJson(&request, w, r) {
		return
	}
	err := InviteMembers(handler.User(r), request)
	if err != nil {
		httpextensions.WriteError(w, err)
		return
	}
	httpextensions.WriteNoContent(w)
}

func SetHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var team = models.Team{}
	if !httpextensions.ReadJson(&team, w, r) {
		return
	}
	team, err := Set(team, handler.User(r))
	if err != nil {
		httpextensions.WriteError(w, err)
		return
	}
	httpextensions.WriteJson(w, team)
}
