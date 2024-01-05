package helpers

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// jwToken is a JWT Token claims struct.
type jwToken struct {
	jwt.StandardClaims
	Iat uint  `json:"iat"`
	Nbf uint  `json:"nbf"`
	ID  uint  `json:"id"`
	Exp int64 `json:"exp"`
}

// GenerateToken generates a JWT token for an id.
func GenerateToken(id uint, jwtSigningKey string, jwtExpirationTime int) string {
	issuedAtTime := uint(time.Now().Unix())
	notBeforeTime := uint(time.Now().Add(time.Minute * -5).Unix())
	expirationTime := time.Now().Add(time.Minute * time.Duration(jwtExpirationTime)).Unix()

	if id == 0 {
		issuedAtTime = uint(time.Now().Add(time.Minute * -15).Unix())
		notBeforeTime = uint(time.Now().Add(time.Minute * -15).Unix())
		expirationTime = time.Now().Add(time.Minute * -15).Unix()
	}

	// Create JWT token
	locJWToken := &jwToken{
		Iat: issuedAtTime,
		Nbf: notBeforeTime,
		Exp: expirationTime,
		ID:  id,
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), locJWToken)
	tokenString, _ := token.SignedString([]byte(jwtSigningKey))

	return tokenString
}
