package jwtUtils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/roverr/go-auth/config"
	"github.com/roverr/go-auth/core/auth/types"
)

// CreateToken can create valid JWT webtokens
// for the client side
func CreateToken(id uint, userName string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS512)
	token.Claims["id"] = id
	token.Claims["userName"] = userName
	expTime := time.Minute * time.Duration(configuration.Conf.JwtExpTime)
	token.Claims["exp"] = time.Now().Add(expTime)
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
	exp, pErr := time.Parse(time.RFC3339, token.Claims["exp"].(string))
	if pErr != nil {
		return authTypes.ParsedToken{}, pErr
	}
	parsedToken := authTypes.ParsedToken{
		UserName: token.Claims["userName"].(string),
		ID:       uint(token.Claims["id"].(float64)),
		Exp:      exp,
	}
	return parsedToken, nil
}
