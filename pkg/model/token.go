package model

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"os"
)

//JWT claims struct
type Token struct {
	UserId uint
	jwt.StandardClaims
}

//Validate user id value
func (t *Token) validate() error {
	if t.UserId <= 0 {
		return fmt.Errorf("Token user ID is empty")
	} else {
		return nil
	}
}

//Create new token and return it's string representation
func (t *Token) NewToken() (string, error) {
	if err := t.validate(); err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), t)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

//Decode token, from Token.NewToken
func (t *Token) DecodeToken(tk string) (uint, error) {
	token, err := jwt.ParseWithClaims(tk, t, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil { //Malformed token, returns with http code 403 as usual
		return 0, err
	}
	if !token.Valid { //Token is invalid, maybe not signed on this server
		return 0, fmt.Errorf("Token is not valid")
	}
	return t.UserId, nil
}