package email

import (
	"logicful/models"
	"logicful/service/httpextensions"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func WebhookHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var header = r.Header.Get("X-Postmark-Webhook")
	if header != os.Getenv("POSTMARK_WEBHOOK_AUTH") {
		httpextensions.WriteNoContent(w)
		return
	}
	var data models.EmailWebhookData
	if !httpextensions.ReadJson(&data, w, r) {
		return
	}
	if data.RecordType == "Bounce" {
		OnBounce(data.Email)
	}
	if data.RecordType == "SubscriptionChange" && data.SuppressSending {
		SetSuppressed(data.Recipient)
	}
	httpextensions.WriteNoContent(w)
}
