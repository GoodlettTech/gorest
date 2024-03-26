package AuthService

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type jwtClaim struct {
	jwt.StandardClaims
	UserID int
}

// CreateToken generates a JWT token for the given user ID.
// It takes the user ID as a parameter and returns the generated token as a string.
// If an error occurs during token generation, it is also returned.
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

// ValidateToken validates the given JWT token string.
// It parses the token using the JWT_SECRET environment variable as the key.
// If the token is valid, it returns the parsed token.
// If the token is invalid, it returns an error indicating the reason.
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

// CreateCookie creates a new HTTP cookie with the given token string.
// The cookie is named "token" and has an expiration time of 1 hour.
func CreateCookie(tokenString string) *http.Cookie {
	return &http.Cookie{
		Name:     "auth",
		Value:    tokenString,
		Expires:  time.Now().Add(time.Hour * 1),
		HttpOnly: true,
		Secure:   true,
	}
}
