package jwtUtils

import (
	"errors"
	"go-auth/auth/types"
	"go-auth/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// CreateToken can create valid JWT webtokens
// for the client side
func CreateToken(id uint, userName string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS512)
	token.Claims["id"] = id
	token.Claims["userName"] = userName
	token.Claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	// tokenString gets []byte because of the HS512 alg
	tokenString, err := token.SignedString([]byte(configuration.Conf.JwtSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func algValidator(token *jwt.Token) (interface{}, error) {
	_, ok := token.Method.(*jwt.SigningMethodHMAC)
	if !ok {
		return nil, errors.New("Signing method failure in JWT token.")
	}
	// Should return []byte since the signing method is HMAC
	return []byte(configuration.Conf.JwtSecret), nil
}

// ValidateToken is a key function for validating
// the JWT tokens sent back from client side
func ValidateToken(tokenString string) (authTypes.ParsedToken, error) {
	token, err := jwt.Parse(tokenString, algValidator)
	if err != nil {
		return authTypes.ParsedToken{}, err
	}
	parsedToken := authTypes.ParsedToken{
		UserName: token.Claims["userName"].(string),
		ID:       token.Claims["id"].(uint),
	}
	return parsedToken, nil
}
