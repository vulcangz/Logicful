package email

import (
	"api/features/user"
	"cloud.google.com/go/firestore"
	"context"
	"github.com/logicful/service/db"
)

func SetSuppressed(email string) {
	u, err := user.ByEmail(email)
	if err != nil {
		return
	}
	_, err = db.Instance().Collection("users").Doc(u.Id).Set(context.Background(), map[string]interface{}{
		"EmailSuppressed": true,
	}, firestore.MergeAll)
	if err != nil {
		return
	}
}

func OnBounce(email string) {
	u, err := user.ByEmail(email)
	if err != nil {
		return
	}
	_, err = db.Instance().Collection("users").Doc(u.Id).Set(context.Background(), map[string]interface{}{
		"EmailBounceCount": firestore.Increment(1),
	}, firestore.MergeAll)
	if err != nil {
		return
	}
}
