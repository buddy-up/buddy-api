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
	ClientID:     os.Getenv("GOOGLE_CLIENT_ID"), 						//The client ID for google oauth. Obtained from google API console
	ClientSecret:  os.Getenv("GOOGLE_CLIENT_SECRET"),					//The secret for google oauth. https://developers.google.com/identity/protocols/OAuth2
	Scopes:       []string{"https://www.googleapis.com/auth/plus.login"},   //The scope you want to use. https://developers.google.com/identity/protocols/googlescopes
	Endpoint:     google.Endpoint,											//Set with the oauth2 library
	RedirectURL:   os.Getenv("REDIRECT_URL"),							//Redirect url. Set using google API console
}

type AccessTokenData struct {
	AccessToken *oauth2.Token	//The access token for the user
	AuthCodeUrl string			//The authentication URL for the user
	Me map[string]interface{}	//The name of the authenticated user
}

/*
	GetAccessToken
	accepts reference to an oauth2 token, and reference to a user object
	Gets the access token for the user
*/
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

	user.Id = me["sub"].(string)

	if getUserData(me, user).Firstname == "" {
		createUser(me)
	}
	return AccessTokenData{accessToken, GOOGLE.AuthCodeURL("state", oauth2.AccessTypeOffline), me}
}

/*
	Authenticate
	Exchanges the token with google oauth
*/
func Authenticate(code string) *oauth2.Token{
	tok, err := GOOGLE.Exchange(oauth2.NoContext, code)
	if err != nil {
		revel.ERROR.Println(err)
		return tok
	}
	return tok
}