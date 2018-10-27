package controllers

import (
	"fmt"
	"os"
	"strconv"

	//"github.com/gomodule/redigo/redis"
	"github.com/revel/revel"
	"github.com/skylerjaneclark/buddy-api/app/models"
	"gopkg.in/redis.v3"
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

	locationSetReply := client.GeoAdd("user_locations", &redis.GeoLocation{Latitude:latitude, Longitude:longitude, Name:user.Id.String()})
	fmt.Println("GET ", locationSetReply)

	fmt.Println(latitude)
	fmt.Println(longitude)
	return c.Redirect(Application.Index)
}

func (c Application) FindNearby (code string) revel.Result {
	user := c.ViewArgs["user"].(*models.User)
	client := RedisConnect()

	res, err := client.GeoRadiusByMember("user_locations", user.Id.String(), &redis.GeoRadiusQuery{
		Radius:      20,
		Unit:        "km",
		WithCoord:   true,
		WithDist:    true,
		Count:       10,
		Sort:        "ASC",
	}).Result()

	if(err != nil){
		fmt.Println(err)
	}
	fmt.Println(res)
	return c.Redirect(Application.Index)

}