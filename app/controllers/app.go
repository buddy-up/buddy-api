package controllers

import (
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/revel/revel"
	"github.com/skylerjaneclark/buddy-api/app/models"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"os"
	"strconv"
)

type Application struct {
	*revel.Controller
}

var GOOGLE = &oauth2.Config{
	ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
	ClientSecret:  os.Getenv("GOOGLE_CLIENT_SECRET"),
	Scopes:       []string{"https://www.googleapis.com/auth/plus.login"},
	Endpoint:     google.Endpoint,
	RedirectURL:   os.Getenv("REDIRECT_URL"),
}

var DB_CONFIG = map[string]string{
	"host" :os.Getenv("DB_HOSTNAME"),
	"port" : "5432",
	"user" : os.Getenv("DB_USER"),
	"password" : os.Getenv("DB_PASSWORD"),
	"dbname" : os.Getenv("DB_NAME"),
}



func (c Application) Index() revel.Result {
	u := c.connected()
	me := make(map[string]interface{})

	if u != nil && u.AccessToken != nil  {
		client :=GOOGLE.Client(oauth2.NoContext, u.AccessToken )
		data, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
		if err != nil {
			fmt.Println("shit")}
		defer data.Body.Close()

		response, _ := ioutil.ReadAll(data.Body)
		json.Unmarshal(response, &me)

		if getUserData(me, c.ViewArgs["user"].(*models.User)).Firstname == "" {
			createUser(me)
		}
			fmt.Println(me)
	}
	authUrl := GOOGLE.AuthCodeURL("state", oauth2.AccessTypeOffline)
	return c.Render(me, authUrl)
}

func (c Application) Auth(code string) revel.Result {
	tok, err := GOOGLE.Exchange(oauth2.NoContext, code)
	if err != nil {
		revel.ERROR.Println(err)
		return c.Redirect(Application.Index)
	}

	user := c.connected()
	user.AccessToken = tok
	return c.Redirect(Application.Index)
}

func (c Application) Logout (code string) revel.Result {
	c.connected().AccessToken = nil
	return c.Redirect(Application.Index)
}

func setuser(c *revel.Controller) revel.Result {
	var user *models.User
	if _, ok := c.Session["uid"]; ok {
		uid, _ := strconv.ParseInt(c.Session["uid"], 10, 0)
		user = models.GetUser(int(uid))
	}
	if user == nil {
		user = models.NewUser()
		c.Session["uid"] = fmt.Sprintf("%d", user.Uid)
	}

	c.ViewArgs["user"] = user
	return nil
}


func init() {
	revel.InterceptFunc(setuser, revel.BEFORE, &Application{})
}

func (c Application) connected() *models.User {
	return c.ViewArgs["user"].(*models.User)
}
