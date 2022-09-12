package user

import (
	"context"
	"errors"
	"fmt"
	"log"
	"logicful/models"
	"logicful/service/date"
	"logicful/service/db"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/api/iterator"
)

func Login(user models.User) (models.TokenResponse, error) {
	byEmail, err := ByEmail(user.Email)
	if err != nil {
		return models.TokenResponse{}, err
	}
	if byEmail.Id == "" || byEmail.Password == "" || !compare(byEmail.Password, []byte(user.Password)) {
		return models.TokenResponse{}, errors.New("invalid username or password")
	}
	token, err := signToken(byEmail)
	if err != nil {
		return models.TokenResponse{}, err
	}
	return token, nil
}

func RefreshToken(token string) models.TokenResponse {
	user := ByToken(token)
	if user.Id == "" {
		return models.TokenResponse{}
	}
	signed, err := signToken(user)
	if err != nil {
		return models.TokenResponse{}
	}
	return signed
}

func ByToken(token string) models.User {
	if token == "" {
		return models.User{}
	}
	var user = models.UserLoginClaims{}
	result, err := jwt.ParseWithClaims(token, &user, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SIGN_TOKEN")), nil
	})

	if err != nil {
		return models.User{}
	}

	if !result.Valid {
		return models.User{}
	}

	return models.User{
		Id:          user.Id,
		Email:       user.Email,
		FullName:    user.FullName,
		DisplayName: user.DisplayName,
		TeamId:      user.TeamId,
		Creatable:   user.Creatable,
	}
}

func ByEmail(email string) (models.User, error) {
	iter := db.Instance().Collection("users").Where("Email", "==", email).Limit(1).Documents(context.Background())
	user := models.User{}
	err := db.First(&user, iter)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func ById(id string) (models.User, error) {
	result, err := db.Instance().Collection("users").Doc(id).Get(context.Background())
	if err != nil {
		return models.User{}, err
	}
	user := models.User{}
	err = db.Unmarshal(result.Data(), &user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func ByTeam(user models.User) ([]models.User, error) {
	iter := db.Instance().Collection("users").Where("TeamId", "==", user.TeamId).Documents(context.Background())
	var results = make([]models.User, 0)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		result := models.User{}
		err = db.Unmarshal(doc.Data(), &result)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}

func Register(user models.User, isSocialLogin bool) (models.TokenResponse, error) {

	user = models.User{
		Email:       user.Email,
		FullName:    user.FullName,
		DisplayName: user.DisplayName,
		Password:    user.Password,
		TeamId:      user.TeamId,
	}

	user.Email = strings.ToLower(user.Email)
	user.Email = strings.TrimSpace(user.Email)
	user.DisplayName = strings.TrimSpace(user.DisplayName)
	user.FullName = strings.TrimSpace(user.FullName)

	exists, err := ByEmail(user.Email)

	if err != nil {
		return models.TokenResponse{}, errors.New("email already exists")
	}

	if exists.Id != "" {
		return models.TokenResponse{}, errors.New("email already exists")
	}

	if user.TeamId == "" {
		user.TeamId = uuid.New().String()
	}

	if user.Id == "" {
		user.Id = uuid.New().String()
	}

	if isSocialLogin {
		user.Password = ""
	} else {
		user.Password = hash([]byte(user.Password))
	}

	user.CreationDate = date.ISO8601(time.Now())
	user.ChangeDate = date.ISO8601(time.Now())

	batch := db.Instance().Batch()
	userRef := db.Instance().Collection("users").Doc(user.Id)
	indexRef := db.Instance().Collection("indexes").Doc("user:" + user.Email)
	batch.Set(userRef, user)
	batch.Set(indexRef, models.Index{
		Key: user.Id,
	})

	_, err = batch.Commit(context.Background())
	if err != nil {
		log.Printf("An error has occurred: %s", err)
		return models.TokenResponse{}, err
	}

	token, err := signToken(user)

	if err != nil {
		return models.TokenResponse{}, err
	}

	return token, nil
}

func hash(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func compare(hashed string, plain []byte) bool {
	byteHash := []byte(hashed)
	err := bcrypt.CompareHashAndPassword(byteHash, plain)
	if err != nil {
		return false
	}
	return true
}

func signToken(user models.User) (models.TokenResponse, error) {
	expiration := time.Now().UTC().Add(time.Hour * 168)
	claims := models.UserLoginClaims{
		Id:          user.Id,
		Email:       user.Email,
		FullName:    user.FullName,
		DisplayName: user.DisplayName,
		TeamId:      user.TeamId,
		Creatable: models.Creatable{
			CreationDate: user.CreationDate,
		},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiration.Unix(),
			Issuer:    "logicful",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	str, err := token.SignedString([]byte(os.Getenv("JWT_SIGN_TOKEN")))
	if err != nil {
		return models.TokenResponse{}, err
	}
	return models.TokenResponse{
		Token:      str,
		Expiration: expiration.UnixNano() / 1000000,
	}, nil
}
