package services

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/akithepriest/click/database"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET_KEY"))

type Claims struct {
	UserID primitive.ObjectID `json:"user_id"`
	Name string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func CreateJWT(user *database.User) (string, error) {
	claims := &Claims{
		UserID: user.ID,
		Name: user.Name,
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer: "click",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func SetJWTCookie(w http.ResponseWriter, tokenString string) {
	cookie := &http.Cookie{
		Name: "access_token",
		Value: tokenString,
		HttpOnly: true,
		Secure: true,
	}
	http.SetCookie(w, cookie)
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func (token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return token, nil
}