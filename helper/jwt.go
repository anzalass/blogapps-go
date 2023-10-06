package helper

import (
	"echodatabase/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const SECRET_JWT = "123"

func CreateToken(username string, email string, id int) (string, error) {
	claims := jwt.MapClaims{}
	claims["id"] = id
	claims["username"] = username
	claims["email"] = email
	// claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.SECRETT))

}

func RefereshToken(accessToken string) string {
	var claims = jwt.MapClaims{}
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	var token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refresh, err := token.SignedString([]byte("secret"))
	if err != nil {
		return ""
	}

	return refresh
}
