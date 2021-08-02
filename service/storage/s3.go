package storage

import (
	"bytes"
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
	"time"
)

func SetJson(value interface{}, name string, bucket string, acl string) (string, error) {

	serialized, err := json.Marshal(value)

	if err != nil {
		return "", err
	}

	return SetJsonBytes(serialized, name, bucket, acl)
}

func GetUrl(key string, bucket string, fileName string) (string, error) {
	if fileName == "" {
		fileName = key
	}
	svc := s3.New(createSession())
	params := &s3.GetObjectInput{
		Bucket:                     aws.String(bucket),
		Key:                        aws.String(key),
		ResponseContentDisposition: aws.String("attachment;filename=" + fileName),
	}
	req, _ := svc.GetObjectRequest(params)
	url, err := req.Presign(15 * time.Minute)
	return url, err
}

func GenerateUrl(key string, bucket string, contentLength int64) (string, error) {
	svc := s3.New(createSession())
	req, _ := svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket:        aws.String(bucket),
		Key:           aws.String(key),
		ContentLength: aws.Int64(contentLength),
	})
	str, err := req.Presign(15 * time.Minute)
	return str, err
}

func SetJsonBytes(value []byte, name string, bucket string, acl string) (string, error) {

	uploader := s3manager.NewUploader(createSession())

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(name),
		ACL:         aws.String(acl),
		ContentType: aws.String("application/json"),
		Body:        bytes.NewReader(value),
	})

	if err != nil {
		return "", err
	}

	return result.Location, nil
}

func SetBytes(value []byte, name string, bucket string, acl string) (string, error) {

	uploader := s3manager.NewUploader(createSession())

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(name),
		ACL:    aws.String(acl),
		Body:   bytes.NewReader(value),
	})

	if err != nil {
		return "", err
	}

	return result.Location, nil
}

func DownloadToFile(bucket string, key string, path string, name string) (bool, error) {
	downloader := s3manager.NewDownloader(createSession())
	err := os.MkdirAll(path, 0644)
	if err != nil {
		return false, err
	}
	file, err := os.Create(path + "/" + name)
	if err != nil {
		return false, nil
	}
	_, err = downloader.Download(file, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if aerr, ok := err.(awserr.Error); ok {
		switch aerr.Code() {
		case s3.ErrCodeNoSuchKey:
			return false, nil
		default:
			return false, err
		}
	}
	return true, nil
}

func DownloadToBytes(bucket string, key string) ([]byte, error) {
	downloader := s3manager.NewDownloader(createSession())
	buff := &aws.WriteAtBuffer{}
	_, err := downloader.Download(buff, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if aerr, ok := err.(awserr.Error); ok {
		switch aerr.Code() {
		case s3.ErrCodeNoSuchKey:
			return nil, nil
		default:
			return nil, err
		}
	}
	return buff.Bytes(), nil
}

func createSession() *session.Session {
	endpoint := os.Getenv("S3_ENDPOINT")
	// The session the S3 Uploader will use
	conf := &aws.Config{
		Region:      aws.String("us-east-1"),
		Endpoint:    &endpoint,
		Credentials: credentials.NewStaticCredentials(os.Getenv("S3_KEY"), os.Getenv("S3_SECRET"), ""),
	}
	return session.Must(session.NewSession(conf))
}
