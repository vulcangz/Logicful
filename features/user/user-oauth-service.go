package user

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/logicful/models"
	"golang.org/x/net/context"
	"google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
	"os"
)

func LoginFromGoogleToken(token string) (models.TokenResponse, error) {
	info, err := verifyIdToken(token)
	if err != nil {
		return models.TokenResponse{}, err
	}
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte("fake_key_not_needed"), nil
	})
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok {
		familyName := fmt.Sprintf("%v", claims["family_name"])
		givenName := fmt.Sprintf("%v", claims["given_name"])
		picture := fmt.Sprintf("%v", claims["picture"])
		email := info.Email
		return loginOrRegister(models.User{
			Email:       email,
			FullName:    givenName + " " + familyName,
			DisplayName: givenName,
			Picture:     picture,
			Creatable:   models.Creatable{},
			Changeable:  models.Changeable{},
		})
	} else {
		return models.TokenResponse{}, err
	}
}

func loginOrRegister(user models.User) (models.TokenResponse, error) {
	token, err := Register(user, true)
	if err != nil && err.Error() == "email already exists" {
		user, err = ByEmail(user.Email)
		if err != nil {
			return models.TokenResponse{}, err
		}
		token, err = signToken(user)
		if err != nil {
			return models.TokenResponse{}, err
		}
		return token, nil
	}
	if err != nil {
		return models.TokenResponse{}, err
	}
	return token, nil
}

func verifyIdToken(idToken string) (*oauth2.Tokeninfo, error) {
	ctx := context.Background()
	authService, err := oauth2.NewService(ctx, option.WithAPIKey(os.Getenv("GOOGLE_OAUTH_SECRET")))
	if err != nil {
		return nil, err
	}
	tokenInfoCall := authService.Tokeninfo()
	tokenInfoCall.IdToken(idToken)
	tokenInfo, err := tokenInfoCall.Do()
	if err != nil {
		return nil, err
	}
	return tokenInfo, nil
}
