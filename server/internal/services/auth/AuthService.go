package AuthService

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type jwtClaim struct {
	jwt.StandardClaims
	UserID int
}

func CreateToken(userId int) (string, error) {
	claim := jwtClaim{
		UserID: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token received")
	}

	return token, nil
}
