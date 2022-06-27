package api

import (
	"github.com/badprinter/remember/internal/models"
	"github.com/golang-jwt/jwt/v4"
)

func createJWT(u *models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodES512)

	// TODO сделать os.env
	key := []byte("remember")
	return token.SignedString(key)
}
