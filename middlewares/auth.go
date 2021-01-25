package middlewares

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4/middleware"
)

type Payload struct {
	ID   string `json:"ID"`
	Name string `json:"name"`
}
type JwtCustomClaims struct {
	ID   string `json:"ID"`
	Name string `json:"name"`
	jwt.StandardClaims
}

func Config() middleware.JWTConfig {
	return middleware.JWTConfig{
		SigningKey: []byte("7PzkBX5DSCl1bjyOxxRbrjUAiQrnQM96"),
		Claims:     &JwtCustomClaims{},
	}
}

func Sign(payload Payload) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["ID"] = payload.ID
	claims["name"] = payload.Name
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, err := token.SignedString([]byte("7PzkBX5DSCl1bjyOxxRbrjUAiQrnQM96"))
	if err != nil {
		return "", err
	}
	return t, nil
}
