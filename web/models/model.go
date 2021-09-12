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

type Tweet struct {
	  UserId uint64        `json:"user_id"`
	  Content string       `json:"content"`
}

type Friendship struct {
	FromId uint64        `json:"from_id"`
	ToId    uint64       `json:"to_id"`
}

type Newsfeed struct {

}