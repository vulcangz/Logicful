package integrations

import (
	"encoding/json"
	"logicful/models"
	"logicful/service/emailer"
	"logicful/service/queue"
	"os"
	"strings"
	"time"

	"github.com/antonholmquist/jason"
)

func RegisterEmail() {
	queue.Receive("send-submission-email", func(message []byte) error {
		var result models.Integration
		err := json.Unmarshal(message, &result)
		if err != nil {
			return err
		}
		return sendSubmissionEmail(result)
	})
}

func sendSubmissionEmail(integration models.Integration) error {
	body, err := formatEmail(integration)
	if err != nil {
		return err
	}
	var sendEmails = os.Getenv("SEND_EMAILS")
	if strings.ToLower(sendEmails) == "true" {
		to := integration.Config["email"]
		test := os.Getenv("TEST_EMAIL_RECEIVER")
		if test != "" {
			to = test
		}
		println("sending_submission_email to" + to)
		var domain = os.Getenv("ROOT_DOMAIN")
		err = emailer.Send(emailer.Email{
			To:       to,
			From:     "maddox@logicful.org",
			Template: "new-submission",
			Model: map[string]string{
				"body":               body,
				"formName":           integration.Form.Title,
				"unsubscribeUrl":     domain + "/unsubscribe",
				"viewSubmissionsUrl": domain + "/form/submissions?formId=" + integration.Submission.FormId,
				"viewSubmissionUrl":  domain + "/form/submissions?formId=" + integration.Submission.FormId + "&submissionId=" + integration.Submission.Id,
			},
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func formatEmail(integration models.Integration) (string, error) {
	var result map[string]interface{}
	var serialized, err = json.Marshal(integration.Submission.Details)
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(serialized, &result)
	if err != nil {
		return "", err
	}

	fields := make([]models.FormField, 0)
	for _, field := range integration.Form.Fields {
		var f models.FormField
		bytes, _ := json.Marshal(field)
		json.Unmarshal(bytes, &f)
		fields = append(fields, f)
	}

	var builder strings.Builder
	for key := range result {
		for _, field := range fields {
			if field.Name == key || field.Label == key {
				value := formatField(field, result[key])
				if value != "" {
					builder.WriteString("<h4><strong>" + field.Label + "</strong></h4>")
					builder.WriteString(value)
					builder.WriteString("<br>")
				}
			}
		}
	}
	var b = builder.String()
	return b, nil
}

func formatField(field models.FormField, data interface{}) string {
	if w, ok := data.(string); ok {
		return "<p>" + w + "</p>"
	}
	if w, ok := data.(float64); ok {
		if field.Type == "date" {
			t := time.Unix(0, int64(w*float64(time.Millisecond)))
			return "<p>" + t.Format(time.RFC822) + "</p>"
		}
	}
	serialized, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	v, err := jason.NewObjectFromBytes(serialized)
	if err != nil {
		return ""
	}
	var builder strings.Builder
	if field.Type == "address" {
		writeField(&builder, v, "Address", "address1", "value")
		writeField(&builder, v, "Address 2", "address2", "value")
		writeField(&builder, v, "City", "city", "value")
		writeField(&builder, v, "State", "state", "value")
		writeField(&builder, v, "Zip Code", "zip", "value")
		return builder.String()
	}
	if field.Type == "date" {
		writeField(&builder, v, "", "formatted")
		return builder.String()
	}
	if field.Type == "full-name" {
		writeField(&builder, v, "Prefix", "prefix", "value")
		writeField(&builder, v, "First", "first", "value")
		writeField(&builder, v, "Middle", "middle", "value")
		writeField(&builder, v, "Last", "last", "value")
		writeField(&builder, v, "Suffix", "suffix", "value")
		return builder.String()
	}
	if field.Type == "file" {
		writeField(&builder, v, "Name", "name")
		writeField(&builder, v, "Type", "type")
		builder.WriteString("<p><strong>" + "Download" + ": " + "</strong>")
		builder.WriteString("View the submission on your submissions page to download the file.")
		builder.WriteString("</p>")
		return builder.String()
	}
	if field.Type == "radio-group" || field.Type == "checkbox-group" {
		radio, err := v.GetObject()
		if err != nil {
			return ""
		}
		builder.WriteString("<p>")
		var values []string
		for _, value := range radio.Map() {
			str, err := value.String()
			if err != nil {
				continue
			}
			values = append(values, str)
		}
		builder.WriteString(strings.Join(values, ", "))
		builder.WriteString("</p>")
		return builder.String()
	}

	builder.WriteString(v.String())
	return builder.String()
}

func writeField(builder *strings.Builder, v *jason.Object, name string, path ...string) {
	val, err := v.GetString(path...)
	if err == nil {
		if name != "" {
			builder.WriteString("<p><strong>" + name + ": " + "</strong>")
		}
		builder.WriteString(val)
		if name != "" {
			builder.WriteString("</p>")
		}
	}
}
