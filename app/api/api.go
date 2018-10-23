package api

import (
	"encoding/json"
	"github.com/revel/revel"
	"github.com/skylerjaneclark/buddy-api/app/models"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"os"
)

var GOOGLE = &oauth2.Config{
	ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
	ClientSecret:  os.Getenv("GOOGLE_CLIENT_SECRET"),
	Scopes:       []string{"https://www.googleapis.com/auth/plus.login"},
	Endpoint:     google.Endpoint,
	RedirectURL:   os.Getenv("REDIRECT_URL"),
}

type AccessTokenData struct {
	AccessToken *oauth2.Token
	AuthCodeUrl string
	Me map[string]interface{}
}

func GetAccessToken (accessToken *oauth2.Token, user *models.User) AccessTokenData{
	me := make(map[string]interface{})

	client :=GOOGLE.Client(oauth2.NoContext, accessToken )
	data, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
	return AccessTokenData{nil, GOOGLE.AuthCodeURL("state", oauth2.AccessTypeOffline), nil}
	}
	defer data.Body.Close()
	response, _ := ioutil.ReadAll(data.Body)
	json.Unmarshal(response, &me)

	if getUserData(me, user).Firstname == "" {
		createUser(me)
	}
	return AccessTokenData{accessToken, GOOGLE.AuthCodeURL("state", oauth2.AccessTypeOffline), me}
}

func Authenticate(code string) *oauth2.Token{
	tok, err := GOOGLE.Exchange(oauth2.NoContext, code)
	if err != nil {
		revel.ERROR.Println(err)
		return tok
	}
	return tok
}