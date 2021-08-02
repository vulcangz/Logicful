package emailer

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

type Email struct {
	To       string
	From     string
	Template string            `json:"TemplateAlias"`
	Model    map[string]string `json:"TemplateModel"`
}

func Send(email Email) error {
	url := "https://api.postmarkapp.com/email/withTemplate"
	payload, err := json.Marshal(email)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Postmark-Server-Token", "0debd00f-760b-415b-97cd-3e80f0d6b905")
	var client = &http.Client{
		Timeout: time.Second * 10,
	}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	println(res.StatusCode)
	println(string(body))
	if res.StatusCode > 299 {
		return errors.New(string(body))
	}
	return nil
}
