package controllers

import (
	"fmt"
	"github.com/revel/revel"
	"github.com/skylerjaneclark/buddy-api/app/api"
)

func (c Application) SaveInstanceId (code string) revel.Result {
	instanceId := c.Params.Form.Get("instanceId")
	user := c.connected()
	user.FireBaseInstanceIds = append(user.FireBaseInstanceIds, instanceId)
	api.StoreInstanceId(*user, instanceId)
	fmt.Println(instanceId)
	fmt.Print("instanceID saved")
	return c.Redirect(Application.Index)
}