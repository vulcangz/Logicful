package optionset

import (
	"context"
	"errors"
	"logicful/models"
	"logicful/service/date"
	"logicful/service/db"
	"time"

	"github.com/google/uuid"
	"google.golang.org/api/iterator"
)

var instance = db.Instance()

func Set(set models.OptionSet, user models.User) error {

	if set.Id == "" {
		set.Id = uuid.New().String()
		set.TeamId = user.TeamId
	}

	if set.TeamId != user.TeamId {
		return errors.New("you do not have access to this option set")
	}

	if set.CreationDate == "" {
		set.CreationDate = date.ISO8601(time.Now())
		set.CreateBy = user.DisplayName
	}

	set.ChangeDate = date.ISO8601(time.Now())
	set.ChangeBy = user.DisplayName

	_, err := instance.Collection("option_sets").Doc(set.Id).Set(context.Background(), set)

	if err != nil {
		return err
	}

	return nil
}

func List(team string) ([]models.OptionSet, error) {
	if team == "" {
		return nil, errors.New("team is required")
	}
	iter := db.Instance().Collection("option_sets").Where("TeamId", "==", team).Documents(context.Background())
	var results = make([]models.OptionSet, 0)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		result := models.OptionSet{}
		err = db.Unmarshal(doc.Data(), &result)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}
