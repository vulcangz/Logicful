package formsubmission

import (
	"api/features/form"
	"cloud.google.com/go/firestore"
	"context"
	"errors"
	"github.com/logicful/models"
	"github.com/logicful/service/db"
)

func Delete(ids []string, formId string, user models.User) error {

	f, err := form.Get(formId)

	if err != nil {
		return err
	}

	if f.TeamId != user.TeamId {
		return errors.New("you do not have access to this form")
	}

	if len(ids) == 0 {
		return nil
	}

	if len(ids) > 450 {
		return errors.New("you may only delete up to 450 entries in a single request")
	}

	var instance = db.Instance()
	var batch = instance.Batch()
	for i := range ids {
		iter := instance.Collection("submissions").Where("FormId", "==", formId).Where("Id", "==", ids[i]).Limit(1).Documents(context.Background())
		doc, err := iter.Next()
		if err != nil {
			continue
		}
		batch.Set(doc.Ref, map[string]interface{}{
			"Status": "deleted",
		}, firestore.MergeAll)
	}
	_, err = batch.Commit(context.Background())
	if err != nil {
		return err
	}
	return nil
}
