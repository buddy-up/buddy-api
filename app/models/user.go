package models

import (
	"golang.org/x/oauth2"
	"math/big"
	"math/rand"
)

type User struct {
	Uid         int				//the UID of the user. Never changes.
	AccessToken *oauth2.Token	//AccessToken of the user
	Firstname string			//The first name of the user
	Lastname string				//The last name of the user
	Email string				//The Email of the user
	Id big.Int					//The ID of the user. Different from UID. UID is saved as a cookie in session. ID is gotten from oauth with their name.
	FireBaseInstanceIds fireBaseIds	//The instanceIDs used by FireBase to send notifications to the user
}

type fireBaseIds struct {
	Web string
	IOS string
	Android string
}

var db = make(map[int]*User)

/*
	GetUser
	Gets the user out of the "db". Used by the session.
*/
func GetUser(id int) *User {
	return db[id]
}

/*
	NewUser
	creates a new user to be used by the session.
*/
func NewUser() *User {
	user := &User{Uid: rand.Intn(10000)}
	db[user.Uid] = user
	return user
}
