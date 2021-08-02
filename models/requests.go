package models

type ChangeFormFolderRequest struct {
	Id           string `json:"id"`
	LastFolderId string `json:"lastFolderId"`
	NewFolderId  string `json:"newFolderId"`
}

type TokenRequest struct {
	Token string `json:"token"`
}

type RemoveTeammatesRequest struct {
	UserIds []string `json:"userIds"`
}

type TeamInviteRequest struct {
	Team Team `json:"team"`
	User User `json:"user"`
	Sender User `json:"sender"`
}
