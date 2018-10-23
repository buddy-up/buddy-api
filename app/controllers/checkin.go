package controllers

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/revel/revel"
	"github.com/skylerjaneclark/buddy-api/app/models"
	"os"
)

func RedisConnect() redis.Conn {
	c, err := redis.Dial(os.Getenv("REDIS_URI"), ":6379")
	if err != nil{
		panic(err)
	}
	return c
}

func (c Application) CheckIn (code string) revel.Result {
	latitude := c.Params.Form.Get("latitude")
	longitude := c.Params.Form.Get("longitude")
	user := c.ViewArgs["user"].(*models.User)

	conn := RedisConnect()
	latitudeReply, latitudeErr := conn.Do("SET", user.Id.String() + "latitude:" ,  latitude)
	longitudeReply, longitudeErr := conn.Do("SET", user.Id.String() +"longitude:" , longitude)
	if latitudeErr != nil{
		panic(latitudeErr)
	}
	if longitudeErr != nil{
		panic(longitudeErr)
	}
	fmt.Println("GET ", longitudeReply)
	fmt.Println("GET ", latitudeReply)
	defer conn.Close()

	fmt.Println(latitude)
	fmt.Println(longitude)
	return c.Redirect(Application.Index)
}