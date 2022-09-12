package integrations

import (
	"bytes"
	"encoding/json"
	"errors"
	"logicful/models"
	"logicful/service/queue"
	"net/http"
	"time"
)

func RegisterWebhook() {
	queue.Receive("send-submission-webhook", func(message []byte) error {
		var result models.Integration
		err := json.Unmarshal(message, &result)
		if err != nil {
			return err
		}
		return sendSubmissionWebhook(result)
	})
}

func sendSubmissionWebhook(integration models.Integration) error {
	url := integration.Config["url"]
	token := integration.Config["token"]
	payload, err := json.Marshal(integration.Submission)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	if token != "" {
		req.Header.Set("X-Logicful-Token", token)
	}
	var client = &http.Client{
		Timeout: time.Second * 10,
	}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode > 299 {
		return errors.New("did not receive success status code on webhook request")
	}
	return nil
}
