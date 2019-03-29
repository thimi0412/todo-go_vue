package main

import (
	"log"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func passwordHash(pw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}

func passwordVerify(hash, pw string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
}

// JWTToken : JWT用のstruct
type JWTToken struct {
	UserID    int    `json:"user_id"`
	UserEmail string `json:"email"`
	jwt.StandardClaims
}

func createTokenString(user User) string {
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), &JWTToken{
		UserID:    user.ID,
		UserEmail: user.Email,
	})

	tokenstring, err := token.SignedString([]byte("foobar"))
	if err != nil {
		log.Fatalln(err)
	}
	return tokenstring
}

func authTokenString(tokenstring string) (int, error) {

	token, err := jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
		return []byte("foobar"), nil
	})

	if err == nil {
		claims := token.Claims.(jwt.MapClaims)

		return int(claims["user_id"].(float64)), err
	}
	return 0, err
}
