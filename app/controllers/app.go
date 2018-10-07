package controllers

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/revel/revel"
	"github.com/skylerjaneclark/buddy-api/app/api"
	"github.com/skylerjaneclark/buddy-api/app/models"
	"strconv"
)

type Application struct {
	*revel.Controller
}

func (c Application) Index() revel.Result {
	u := c.connected()
	var tokenData api.AccessTokenData
	tokenData = api.GetAccessToken(u.AccessToken, c.ViewArgs["user"].(*models.User))
	u.AccessToken = tokenData.AccessToken

	var me = tokenData.Me
	var authUrl = tokenData.AuthCodeUrl
	return c.Render(me, authUrl)
}

func (c Application) Auth(code string) revel.Result {
	var tok = api.Authenticate(code)
	user := c.connected()
	user.AccessToken = tok
	return c.Redirect(Application.Index)
}

func (c Application) Logout (code string) revel.Result {
	c.connected().AccessToken = nil
	return c.Redirect(Application.Index)
}


func (c Application) Location (location map[string]interface{}) revel.Result {
	return nil
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
