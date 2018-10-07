package api


func setLocation(location  map[string]interface{}){
	//basically, go to redis, set the key of userID:location:whatever to their current locatio

}

func getLocation ()  {
	//gets the userID:location and returns where they are

}

func pingLocation()  {
	//sends out the request every 20 minutes for people to update location

}

func grabLocation(){
	//sends out an immediate request of closest people

}