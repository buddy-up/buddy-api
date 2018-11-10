package controllers

import (
	"firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"fmt"
	"github.com/revel/revel"
	"github.com/skylerjaneclark/buddy-api/app/api"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"gopkg.in/redis.v3"
	"os"
	"strconv"
)

/*
	RedisConnect
	connects to the redis DB
*/
func RedisConnect() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:    os.Getenv("REDIS_URI") + os.Getenv("REDIS_PORT")  ,
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,  // use default DB
	})
	return client
}

/*
	CheckIn
	Saves the location of a user.
	Makes a connection to redis, and sets the latitude and longitude of a user before redirecting back to the main app.
*/
func (c Application) CheckIn (code string) revel.Result {
	latitude, latitudeErr := strconv.ParseFloat(c.Params.Form.Get("latitude"), 64)
	longitude, longitudeErr := strconv.ParseFloat(c.Params.Form.Get("longitude"), 64)
	if longitudeErr != nil {
		fmt.Println(longitudeErr)
	}
	if latitudeErr != nil {
		fmt.Println(latitudeErr)
	}
	user := c.connected()
	client := RedisConnect()

	if user.FireBaseInstanceIds.Android != "" {
		locationSetReply := client.GeoAdd("user_locations", &redis.GeoLocation{Latitude:latitude, Longitude:longitude, Name:user.FireBaseInstanceIds.Android})
		fmt.Println("GET ", locationSetReply)
	}else if user.FireBaseInstanceIds.IOS != ""{
		locationSetReply := client.GeoAdd("user_locations", &redis.GeoLocation{Latitude:latitude, Longitude:longitude, Name:user.FireBaseInstanceIds.IOS})
		fmt.Println("GET ", locationSetReply)
	}else{
		locationSetReply := client.GeoAdd("user_locations", &redis.GeoLocation{Latitude:latitude, Longitude:longitude, Name:user.FireBaseInstanceIds.Web})
		fmt.Println("GET ", locationSetReply)
	}

	fmt.Println(latitude)
	fmt.Println(longitude)
	return c.Redirect(Application.Index)
}

func (c Application) RemoveGeoLocation() revel.Result{
	instanceId := c.Params.Form.Get("instanceId")

	client := RedisConnect()
	res, err := client.ZRem("user_locations", instanceId).Result()
	fmt.Println(res)
	fmt.Println(err)
	return c.Redirect(Application.Index)
}

func (c Application) FindNearby (code string) revel.Result {
	user := c.connected()
	client := RedisConnect()

	api.GetInstanceIds(user)
	userId := ""
	if user.FireBaseInstanceIds.Android != "" {
		userId = user.FireBaseInstanceIds.Android
	}else if user.FireBaseInstanceIds.IOS != ""{
		userId = user.FireBaseInstanceIds.IOS
	}else if user.FireBaseInstanceIds.Web != ""{
		userId = user.FireBaseInstanceIds.Web
	}

	res, err := client.GeoRadiusByMember("user_locations",userId, &redis.GeoRadiusQuery{
		Radius:      20,
		Unit:        "km",
		WithCoord:   true,
		WithDist:    true,
		Count:       10,
		Sort:        "ASC",
	}).Result()
	
	for index, element := range res {

		opt := option.WithCredentialsFile(os.Getenv("FIREBASE_CONF_LOCATION"))

		app, err := firebase.NewApp(context.Background(), nil,opt)
		if err != nil {
			panic(err)
			return nil
		}
		ctx := context.Background()
		client, err := app.Messaging(ctx)

		registrationToken := element.Name

		message := &messaging.Message{
			Data: map[string]string{
				"score": "850",
				"time":  "2:45",
			},
			Token: registrationToken,
		}
		response, err := client.Send(ctx, message)
		if err != nil {
			panic(err)
			return nil
		}
		fmt.Println("Successfully sent message:", response)
		fmt.Println(index)
	}

		if(err != nil){
		fmt.Println(err)
	}
	fmt.Println(res)
	return c.Redirect(Application.Index)

}