package services

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const GOOGLE_OAUTH_USERINFO = "https://www.googleapis.com/oauth2/v1/userinfo?access_token="

var dataScopes = []string{
	"https://www.googleapis.com/auth/userinfo.email",
	"https://www.googleapis.com/auth/userinfo.profile",
}

type GoogleUserInfo struct {
	ID string `json:"id"`
	Email string `json:"email"`
	VerifiedEmail bool `json:"verified_email"`
	Name string `json:"name"`
	GivenName string `json:"given_name"`
	Picture string `json:"picture"`
}

type GoogleOAuthService struct {
	config *oauth2.Config
}

func NewGoogleOAuthService(clientID string, clientSecret string, callbackUrl string) *GoogleOAuthService {
	config := &oauth2.Config{
		RedirectURL: callbackUrl,
		ClientID: clientID,
		ClientSecret: clientSecret,
		Scopes: dataScopes,
		Endpoint: google.Endpoint,
	}
	return &GoogleOAuthService{
		config: config,
	}
}

func (s *GoogleOAuthService) GetLoginUrl() string {
	state := s.generateStateOauthCookie()
	u := s.config.AuthCodeURL(state)
	return u
}

// Exchange auth code with google for user data.
// Return: name, email, error
func (s *GoogleOAuthService) GetUserData(ctx context.Context, code string) (userinfo *GoogleUserInfo, err error) {
	token, err := s.config.Exchange(ctx, code)
	if err != nil {
		return nil, errors.New("invalid token")
	}

	res, err := http.Get(GOOGLE_OAUTH_USERINFO + token.AccessToken)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	contents, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}

	if err = json.Unmarshal(contents, &userinfo); err != nil {
		return nil , fmt.Errorf("failed to marshal: %v", err)
	}

	return userinfo, nil
}

func (s *GoogleOAuthService) generateStateOauthCookie() string {
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)

	return state
}