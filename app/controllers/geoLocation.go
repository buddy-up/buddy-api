package controllers

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/revel/revel"
	"github.com/skylerjaneclark/buddy-api/app/models"
	"os"
)

/*
	RedisConnect
	connects to the redis DB
*/
func RedisConnect() redis.Conn {
	c, err := redis.Dial(os.Getenv("REDIS_URI"), os.Getenv("REDIS_PORT")) //uses redigo to dial the URI and port
	if err != nil{
		panic(err)
	}
	response, err := c.Do("AUTH", os.Getenv("REDIS_PASSWORD")) //if your redis has a password you need this
	fmt.Printf("Connected! ", response)
	return c
}

/*
	CheckIn
	Saves the location of a user.
	Makes a connection to redis, and sets the latitude and longitude of a user before redirecting back to the main app.
*/
func (c Application) CheckIn (code string) revel.Result {
	latitude := c.Params.Form.Get("latitude")
	longitude := c.Params.Form.Get("longitude")
	user := c.ViewArgs["user"].(*models.User)

	conn := RedisConnect()

	locationSetReply, locationSetErr := conn.Do("GEOADD", "user_locations", latitude, longitude, user.Id)

	if locationSetErr != nil{
		panic(locationSetErr)
	}
	fmt.Println("GET ", locationSetReply)
	defer conn.Close()

	fmt.Println(latitude)
	fmt.Println(longitude)
	return c.Redirect(Application.Index)
}

func (c Application) FindNearby (code string) revel.Result {
	user := c.ViewArgs["user"].(*models.User)

	conn := RedisConnect()

	nums, err := redis.Values(conn.Do("GEORADIUS", user.Id, 15, 37, 200, "km", "WITHCOORD", "WITHDIST"))
	if(err != nil){
		fmt.Print(err)
	}
	fmt.Println(nums)
	return c.Redirect(Application.Index)

}