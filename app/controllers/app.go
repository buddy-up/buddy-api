package controllers

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/revel/revel"
	"github.com/skylerjaneclark/buddy-api/app/api"
	"github.com/skylerjaneclark/buddy-api/app/models"
	"golang.org/x/oauth2"
	"strconv"
)

type Application struct {
	*revel.Controller
}

/*
	Index
	The starting point of the application
	Gets the connection, calls getAccessToken to set the session's accessToken.
	The authCodeUrl is passed to the template and set as the url for the button to log in, which sends the user to the google login page.
	if the tokenData.me hasn't been set, the function renders me with nothing, and that case is handled in the template.
 */
func (c Application) Index() revel.Result {
	u := c.connected()
	var tokenData api.AccessTokenData
	tokenData = api.GetAccessToken(u.AccessToken, c.ViewArgs["user"].(*models.User))
	u.AccessToken = tokenData.AccessToken

	var me = tokenData.Me
	var authUrl = tokenData.AuthCodeUrl
	return c.Render(me, authUrl)
}

/*
	Auth
	sets the session's access token, and redirects to the index. The final step of the Oauth tango.
*/
func (c Application) Auth(code string) revel.Result {
	tok, err := api.GOOGLE.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Println(err)
		return c.Redirect(Application.Index)
	}
	user := c.connected()
	user.AccessToken = tok
	return c.Redirect(Application.Index)}

/*
	Logout
	sets connected user's access token to nil, thus logging them out.
*/
func (c Application) Logout (code string) revel.Result {
	c.connected().AccessToken = nil
	c.ViewArgs["user"] = nil
	return c.Redirect(Application.Index)
}

/*
	setuser
	makes a user, and if the session has a UID then it parses the UID and sets the user's UID to the session UID.
	if the user doesn't exist, it creates a new user and sets the UID to the UID created by the user that was just created.
	After that is sets the user in the session to the user that was either craeted or edited.
*/
func setuser(c *revel.Controller) revel.Result {
	var user *models.User
	if _, ok := c.Session["uid"]; ok {
		uid, _ := strconv.ParseInt(c.Session["uid"].(string), 10, 0)
		user = models.GetUser(int(uid))
	}
	if user == nil {
		user = models.NewUser()
		c.Session["uid"] = fmt.Sprintf("%d", user.Uid)
	}
	c.ViewArgs["user"] = user
	return nil
}
/*
	init
	Intercepts the app running to call the setuser function
*/
func init() {
	revel.InterceptFunc(setuser, revel.BEFORE, &Application{})
}

/*
	connected
	sets the user once the application connects.
*/
func (c Application) connected() *models.User {
	return c.ViewArgs["user"].(*models.User)
}
