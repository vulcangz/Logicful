package db

import (
	"cloud.google.com/go/firestore"
	"encoding/json"
	"google.golang.org/api/iterator"
)

func First(result interface{}, iter *firestore.DocumentIterator) error {
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		err = Unmarshal(doc.Data(), &result)
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}

func Unmarshal(doc map[string]interface{}, result interface{}) error {
	bytes, err := json.Marshal(doc)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return err
	}
	return nil
}
