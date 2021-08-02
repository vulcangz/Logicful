package models

type EmailWebhookData struct {
	RecordType      string
	Recipient       string
	SuppressSending bool
	Email           string
}
