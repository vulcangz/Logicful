package team

import (
	"api/features/user"
	"cloud.google.com/go/firestore"
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/logicful/models"
	"github.com/logicful/service/date"
	"github.com/logicful/service/db"
	"github.com/logicful/service/queue"
	"time"
)

func ByUser(user models.User) (models.Team, error) {
	iter := db.Instance().Collection("teams").Where("Id", "==", user.TeamId).Limit(1).Documents(context.Background())
	team := models.Team{}
	err := db.First(&team, iter)
	if err != nil {
		return models.Team{}, err
	}
	return team, nil
}

func IsOwnerOrAdmin(user models.User, team models.Team) bool {
	if team.Owner == user.Id {
		return true
	}
	for _, admin := range team.Admins {
		if admin == user.Id {
			return true
		}
	}
	return false
}

func InviteMembers(admin models.User, emails []string) error {
	if len(emails) == 0 {
		return nil
	}
	team, err := ByUser(admin)
	if err != nil {
		return err
	}
	if !IsOwnerOrAdmin(admin, team) {
		return errors.New("you do not have permission to add users")
	}
	var instance = db.Instance()
	var batch = instance.Batch()
	var users = make([]models.User, 0)
	for _, email := range emails {
		byEmail, err := user.ByEmail(email)
		if err != nil {
			continue
		}
		if byEmail.TeamId == team.Id {
			continue
		}
		if byEmail.PendingTeamInvites[team.Id] == true {
			continue
		}
		batch.Set(instance.Collection("users").Doc(byEmail.Id), map[string]interface{}{
			"PendingTeamInvites": map[string]interface{}{
				team.Id: true,
			},
		}, firestore.MergeAll)
		users = append(users, byEmail)
	}
	if len(users) == 0 {
		return nil
	}
	for _, u := range users {
		err := queue.Dispatch("team-invite", models.TeamInviteRequest{
			Team:   team,
			User:   u,
			Sender: admin,
		})
		if err != nil {
			return err
		}
	}
	_, err = batch.Commit(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func AcceptInvite(u models.User, teamId string) error {
	full, err := user.ById(u.Id)
	if err != nil {
		return err
	}
	if full.PendingTeamInvites[teamId] == false {
		return errors.New("no invite found by that team id")
	}
	_, err = db.Instance().Collection("users").Doc(u.Id).Set(context.Background(), map[string]interface{}{
		"TeamId": teamId,
		"PendingTeamInvites": map[string]interface{}{
			teamId: false,
		},
	}, firestore.MergeAll)
	if err != nil {
		return err
	}
	return nil
}

func RemoveMembers(admin models.User, userIds []string) error {
	if len(userIds) == 0 {
		return nil
	}
	team, err := ByUser(admin)
	if err != nil {
		return err
	}
	if !IsOwnerOrAdmin(admin, team) {
		return errors.New("you do not have permission to remove users")
	}
	var instance = db.Instance()
	var batch = instance.Batch()
	for i := range userIds {
		if userIds[i] == team.Owner {
			return errors.New("you cannot remove the owner")
		}
		iter := db.Instance().Collection("users").Where("TeamId", "==", admin.TeamId).Where("Id", "==", userIds[i]).Documents(context.Background())
		doc, err := iter.Next()
		if err != nil {
			continue
		}
		batch.Set(doc.Ref, map[string]interface{}{
			"TeamId": userIds[i],
		}, firestore.MergeAll)
	}
	_, err = batch.Commit(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func Set(team models.Team, user models.User) (models.Team, error) {

	team.Id = user.TeamId

	if team.Id == "" {
		team.Id = uuid.New().String()
	}

	if team.CreationDate == "" {
		team.CreationDate = date.ISO8601(time.Now())
		team.CreateBy = user.DisplayName
	}

	team.ChangeDate = date.ISO8601(time.Now())
	team.ChangeBy = user.DisplayName

	_, err := db.Instance().Collection("teams").Doc(team.Id).Set(context.Background(), team)

	if err != nil {
		return models.Team{}, err
	}

	return team, err
}
