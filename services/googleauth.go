package services

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var dataScopes = []string{
	"https://www.googleapis.com/auth/userinfo.email",
	"https://www.googleapis.com/auth/userinfo.profile",
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

func (s *GoogleOAuthService) GetLoginUrl(w http.ResponseWriter) string {
	state := s.generateStateOauthCookie(w)
	u := s.config.AuthCodeURL(state)
	return u
}

func (s *GoogleOAuthService) generateStateOauthCookie(w http.ResponseWriter) string {
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)

	return state
}