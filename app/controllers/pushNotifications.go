package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/revel/revel"
	"github.com/skylerjaneclark/buddy-api/app/api"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type idInfo struct {
application string
subtype string
scope string
authorizedEntity string
platform string
}

func (c Application) SaveInstanceId (code string) revel.Result {
	instanceId := c.Params.Form.Get("instanceId")
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://iid.googleapis.com/iid/info/" + instanceId, nil)
	if err != nil{
		log.Fatal(err)
	}
	req.Header.Add("Authorization", "key=" + os.Getenv("FIREBASE_SERVER_KEY"))
	res, reqErr := client.Do(req)
	if reqErr != nil{
		log.Fatal(reqErr)
	}
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	idInfo := idInfo{}
	jsonErr := json.Unmarshal(body, &idInfo)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	user := c.connected()

	if idInfo.platform == "ANDROID" {
		user.FireBaseInstanceIds.Android = instanceId
	} else if idInfo.platform == "BROWSER"{
		user.FireBaseInstanceIds.Web = instanceId
	} else {
		user.FireBaseInstanceIds.IOS = instanceId
	}

	api.StoreInstanceId(user, instanceId, idInfo.platform)
	fmt.Println(instanceId)
	fmt.Println("instanceID saved")
	return c.Redirect(Application.Index)
}