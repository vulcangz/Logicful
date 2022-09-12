package handler

import (
	"logicful/features/user"
	"logicful/models"
	"net/http"
	"strings"
)

func User(r *http.Request) models.User {
	if r.Header.Get("Authorization") == "" {
		return models.User{}
	}
	var token = r.Header.Get("Authorization")
	token = strings.Replace(token, "Bearer ", "", 1)
	token = strings.TrimSpace(token)
	return user.ByToken(token)
}
