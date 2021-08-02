package formsubmission

import (
	"api/features/form"
	"cloud.google.com/go/firestore"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/logicful/models"
	"github.com/logicful/service/date"
	"github.com/logicful/service/db"
	"github.com/logicful/service/storage"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func MarkRead(ids map[string]bool, formId string, user models.User) error {
	client := db.Instance()
	batch := client.Batch()
	ref := db.Instance().Collection("submissions_seen").Doc(user.Id + ":" + formId)
	for s, b := range ids {
		batch.Set(ref, map[string]interface{}{
			"Submissions": map[string]interface{}{
				s: b,
			},
		}, firestore.MergeAll)
	}
	_, err := batch.Commit(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func GetRead(formId string, user models.User) (map[string]bool, error) {
	doc, err := db.Instance().Collection("submissions_seen").Doc(user.Id + ":" + formId).Get(context.Background())
	if err != nil && status.Code(err) == codes.NotFound {
		return make(map[string]bool, 0), nil
	}
	bytes, err := json.Marshal(doc.Data()["Submissions"])
	if err != nil {
		return nil, err
	}
	var m map[string]bool
	err = json.Unmarshal(bytes, &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func Add(submission models.Submission) error {

	if submission.FormId == "" {
		return errors.New("formId is required on submission")
	}

	if submission.FormId == "demo" {
		return errors.New("demo form may not be submitted")
	}

	submission.Id = uuid.New().String()

	submission.CreationDate = date.ISO8601(time.Now())
	submission.CreateBy = "maddox"
	submission.Processed = false
	submission.ReadBy = make(map[string]bool, 0)

	_, err := db.Instance().Collection("submissions").Doc(submission.Id).Set(context.Background(), submission)

	if err != nil {
		return err
	}

	return nil
}

func List(id string, user models.User) (string, error) {
	f, err := form.Get(id)
	if err != nil {
		return "", err
	}
	if f.TeamId == "" || f.TeamId != user.TeamId {
		return "", errors.New("you do not have access to submissions")
	}
	name := id + ".json"
	url, err := storage.GetUrl(name, "logicful-form-submissions", "")
	if err != nil {
		return "", err
	}
	return url, nil
}

func GetSubmissionFile(file models.File, user models.User) (string, error) {
	f, err := form.Get(file.FormId)
	if err != nil {
		return "", err
	}
	if f.TeamId == "" || f.TeamId != user.TeamId {
		return "", errors.New("you do not have access to submissions")
	}
	url, err := storage.GetUrl(file.Id, "logicful-form-assets", file.Name)
	if err != nil {
		return "", err
	}
	return url, nil
}
