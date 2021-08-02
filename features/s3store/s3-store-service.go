package s3store

import (
	"github.com/google/uuid"
	"github.com/logicful/service/storage"
)

func SetJson(key string, value []byte) (string, error) {
	if key == "" {
		key = uuid.New().String()
	}
	name := key

	path, err := storage.SetJsonBytes(value, name, "json-data", "public-read")

	if err != nil {
		return "", err
	}

	return path, nil
}

func GenerateStoreUrl(contentLength int64) (string, string, error) {
	key := uuid.New().String()
	url, err := storage.GenerateUrl(key, "logicful-form-assets", contentLength)
	if err != nil {
		return "", "", err
	}
	return url, key, nil
}
