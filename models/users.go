package models

import "github.com/dgrijalva/jwt-go"

type User struct {
	Id          string `json:"id" dynamodbav:"userId"`
	Email       string `json:"email"`
	FullName    string `json:"fullName"`
	DisplayName string `json:"displayName"`
	Password    string `json:"password,omitempty"`
	TeamId      string `json:"teamId"`
	Picture     string `json:"picture"`
	Creatable
	Changeable
	EmailSuppressed  string `json:"emailSuppressed"`
	EmailBounceCount string `json:"emailBounceCount"`
	PendingTeamInvites map[string]bool `json:"pendingTeamInvites"`
}

type UserLoginClaims struct {
	Id          string `json:"id"`
	Email       string `json:"email"`
	FullName    string `json:"fullName"`
	DisplayName string `json:"displayName"`
	TeamId      string `json:"teamId"`
	Creatable
	jwt.StandardClaims
}
