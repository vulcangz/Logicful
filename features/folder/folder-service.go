package folder

import (
	"context"
	"errors"
	"logicful/models"
	"logicful/service/date"
	"logicful/service/db"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/google/uuid"
	"google.golang.org/api/iterator"
)

func Set(folder models.Folder, user models.User) (models.Folder, error) {

	if folder.TeamId == "" {
		return models.Folder{}, errors.New("team id is required")
	}

	if folder.TeamId != user.TeamId {
		return models.Folder{}, errors.New("you do not have access to this folder")
	}

	if folder.Id == "" {
		folder.Id = uuid.New().String()
	}

	if folder.CreationDate == "" {
		folder.CreationDate = date.ISO8601(time.Now())
		folder.CreateBy = user.DisplayName
	}

	folder.ChangeDate = date.ISO8601(time.Now())
	folder.ChangeBy = user.DisplayName

	_, err := db.Instance().Collection("folders").Doc(folder.Id).Set(context.Background(), folder)

	if err != nil {
		return models.Folder{}, err
	}

	return folder, err
}

func Delete(id string, user models.User) error {

	children, err := hasChildren(id)
	if err != nil {
		return err
	}

	if children {
		return errors.New("must delete all child folders first")
	}

	client := db.Instance()
	batch := client.Batch()
	ctx := context.Background()

	iter := client.Collection("forms").Where("Folder", "==", id).Where("TeamId", "==", user.TeamId).Documents(ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		batch.Set(doc.Ref, map[string]interface{}{
			"Folder": user.TeamId + ":" + "uncategorized",
		}, firestore.MergeAll)
	}

	batch.Delete(client.Collection("folders").Doc(id))

	_, err = batch.Commit(ctx)

	if err != nil {
		return err
	}
	return nil
}

func List(team string) ([]models.Folder, error) {
	if team == "" {
		return nil, errors.New("team is required")
	}
	iter := db.Instance().Collection("folders").Where("TeamId", "==", team).Documents(context.Background())
	var results = make([]models.Folder, 0)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		result := models.Folder{}
		err = db.Unmarshal(doc.Data(), &result)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}

func hasChildren(folder string) (bool, error) {
	iter := db.Instance().Collection("folders").Where("Parent", "==", folder).Documents(context.Background())
	doc, err := iter.Next()
	if err == iterator.Done {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return doc.Exists(), nil
}
