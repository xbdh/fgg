package models


import "github.com/dgrijalva/jwt-go"

type CustomClaims struct {
	UserId uint64
	NickName    string
	jwt.StandardClaims
}

type PassWordLoginForm struct {
	Name    string        `json:"name"`
	Password  string      `json:"password"`

}



type RegisterForm struct {
	Name    string        `json:"name"`
	Password string       `json:"password"`
	Email string          `json:"email"`
}
