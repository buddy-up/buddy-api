package models

import (
	"golang.org/x/oauth2"
	"math/big"
	"math/rand"
)

type User struct {
	Uid         int
	AccessToken *oauth2.Token
	Firstname string
	Lastname string
	Email string
	Id big.Int
}

var db = make(map[int]*User)

func GetUser(id int) *User {
	return db[id]
}

func NewUser() *User {
	user := &User{Uid: rand.Intn(10000)}
	db[user.Uid] = user
	return user
}
