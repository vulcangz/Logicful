package team

import (
	"encoding/json"
	"github.com/logicful/models"
	"github.com/logicful/service/emailer"
	"github.com/logicful/service/queue"
	"net/url"
	"os"
)

func Register() {
	registerEmailSender()
}

func registerEmailSender() {
	queue.Receive("team-invite-email", func(message []byte) error {
		var result models.TeamInviteRequest
		err := json.Unmarshal(message, &result)
		if err != nil {
			return err
		}
		var domain = os.Getenv("ROOT_DOMAIN")
		err = emailer.Send(emailer.Email{
			To:       result.User.Email,
			From:     "maddox@logicful.org",
			Template: "user-invitation",
			Model: map[string]string{
				"product_url":                     domain,
				"product_name":                    "Logicful Forms",
				"team_name":                       result.Team.Name,
				"name":                            result.User.DisplayName,
				"sender_name":                     result.Sender.DisplayName,
				"action_url":                      domain + "/team/accept?team=" + result.Team.Id + "&name=" + url.QueryEscape(result.Team.Name),
				"help_url":                        "https://docs.logicful.org",
				"company_name":                    "Logicful",
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
