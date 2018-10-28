package controllers

import (
	"fmt"
	"github.com/revel/revel"
	"github.com/skylerjaneclark/buddy-api/app/api"
)

func (c Application) SaveInstanceId (code string) revel.Result {
	instanceId := c.Params.Form.Get("instanceId")
	origin := c.Params.Form.Get("origin")
	user := c.connected()

	if origin == "ANDROID" {
		user.FireBaseInstanceIds.Android = instanceId
	} else if origin == "CHROME"{
		user.FireBaseInstanceIds.Web = instanceId
	} else {
		user.FireBaseInstanceIds.IOS = instanceId
	}
	api.StoreInstanceId(*user, instanceId, origin)
	fmt.Println(instanceId)
	fmt.Println("instanceID saved")
	return c.Redirect(Application.Index)
}