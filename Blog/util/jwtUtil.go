package util

import (
	"Blog/github.com/dgrijalva/jwt-go"
	"fmt"
	"time"
)

type JwtClaims struct {
	//JwtClaims	Verification Jwt And Read Claims
	jwt.StandardClaims // Inheritance of common JWT variables
	Uid string `json:"uid"`
}

const secretKey  = "Cain"

type JwtUtil struct {
	//Jwt Util Struct New And Verification
}



func (JwtUtil) New(uid string)  (string, error) {
	//New Jwt

	claims := JwtClaims{
		Uid:            uid,
	}
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(3600)).Unix() //exp
	claims.IssuedAt = time.Now().Unix() //iat
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))

	if err != nil {
		fmt.Println(err)
		return "", err
	}


	return tokenString, nil
}

func (JwtUtil) 	Verification(token string) (*JwtClaims, bool) {
	//Verification Jwt return JwtClaims
	jwtClaims := &JwtClaims{}
	_, err := jwt.ParseWithClaims(token, jwtClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		fmt.Println("Errer: ", err)
		return jwtClaims, false
	}

	return jwtClaims, true
}
