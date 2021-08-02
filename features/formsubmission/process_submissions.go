package formsubmission

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"github.com/logicful/models"
	"github.com/logicful/service/db"
	"github.com/logicful/service/distributedlock"
	"github.com/logicful/service/queue"
	"github.com/logicful/service/storage"
	"github.com/robfig/cron/v3"
	"google.golang.org/api/iterator"
	"log"
)

func StartProcessor() {
	c := cron.New()
	_, err := c.AddFunc("@every 1m", process)
	if err != nil {
		log.Fatal(err)
	}
	c.Start()
}

func process() {
	println("Checking unprocessed submissions.")
	submissions, err := Unprocessed()
	if err != nil {
		println("Error getting submissions to process " + err.Error())
		return
	}
	if len(submissions) == 0 {
		//println("No submissions to process.")
		return
	}
	formMap := make(map[string][]models.Submission)
	for _, sub := range submissions {
		if sub.FormId == "" || sub.Id == "" {
			continue
		}
		if formMap[sub.FormId] == nil {
			formMap[sub.FormId] = make([]models.Submission, 1)
		}
		formMap[sub.FormId] = append(formMap[sub.FormId], sub)
	}
	for s := range formMap {
		mutex, err := distributedlock.Acquire(s)
		if err != nil {
			println("Error acquiring submissions lock " + err.Error())
			return
		}
		err = ProcessForm(s, formMap[s])
		if err != nil {
			println("Failed to process submission entries " + err.Error())
			_ = distributedlock.Release(mutex)
			return
		}
		_ = distributedlock.Release(mutex)
	}
}

func ProcessForm(formId string, submissions []models.Submission) error {
	name := formId + ".json"
	current, err := currentSubmissions(name)
	if err != nil {
		return err
	}
	for _, submission := range submissions {
		if submission.Status == "deleted" {
			var before = len(current)
			current = removeFromArray(current, submission)
			if (len(current)) == before {
				println("already deleted.")
				continue
			}
		} else {
			// Already exists, just return.
			for i := range current {
				if current[i].Id == submission.Id {
					println("already exists.")
					continue
				}
			}
			current = append(current, submission)
		}
	}

	for _, submission := range current {
		if submission.Id == "" {
			current = removeFromArray(current, submission)
		}
	}

	_, err = storage.SetJson(current, name, "logicful-form-submissions", "private")

	if err != nil {
		return err
	}

	err = SetUnprocessed(submissions)

	if err != nil {
		return err
	}

	return nil
}

func currentSubmissions(name string) ([]models.Submission, error) {
	bytes, err := storage.DownloadToBytes("logicful-form-submissions", name)
	if err != nil {
		return nil, err
	}
	if bytes == nil {
		return make([]models.Submission, 0), nil
	}

	var submissions []models.Submission
	err = json.Unmarshal(bytes, &submissions)

	if err != nil {
		return nil, err
	}

	return submissions, nil
}

func SetUnprocessed(submissions []models.Submission) error {
	client := db.Instance()
	batch := client.Batch()
	formId := ""
	count := 0

	for _, s := range submissions {
		if s.Id == "" {
			continue
		}
		err := queue.Dispatch("submissions", s)
		if err != nil {
			return err
		}
		count++
		formId = s.FormId
		doc := client.Collection("submissions").Doc(s.Id)
		batch.Set(doc, map[string]interface{}{
			"Processed": true,
		}, firestore.MergeAll)
	}

	batch.Update(db.Instance().Collection("forms").Doc(formId), []firestore.Update{
		{Path: "SubmissionCount", Value: firestore.Increment(count)},
		{Path: "UnreadSubmissions", Value: firestore.Increment(count)},
	})

	_, err := batch.Commit(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func Unprocessed() ([]models.Submission, error) {
	iter := db.Instance().Collection("submissions").Where("Processed", "==", false).Limit(450).Documents(context.Background())
	results := make([]models.Submission, 0)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		result := models.Submission{}
		err = db.Unmarshal(doc.Data(), &result)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}

func removeFromArray(s []models.Submission, submission models.Submission) []models.Submission {
	i := 0
	for _, x := range s {
		if x.Id != submission.Id {
			s[i] = x
			i++
		}
	}
	s = s[:i]
	return s
}
