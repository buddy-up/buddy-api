package controllers

import (
	"database/sql"
	"fmt"
)

func createUser(userData map[string]interface{}){
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+ "password=%s dbname=%s",
		DB_CONFIG["host"], DB_CONFIG["port"], DB_CONFIG["user"], DB_CONFIG["password"], DB_CONFIG["dbname"])

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if len(userData) != 0{
		sqlStatement := `
			INSERT INTO users (id, firstname, lastname, email )
			VALUES (1,'`+ userData["given_name"].(string)+`','`+ userData["family_name"].(string)+`','`+ userData["email"].(string)+`');`


		fmt.Printf("______________________________")
		fmt.Println("map:", sqlStatement)
		_, err = db.Exec(sqlStatement)
		if err != nil {
			panic(err)
		}
	}



}