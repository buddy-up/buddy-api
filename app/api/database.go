package api

import (
	"database/sql"
	"fmt"
	"github.com/skylerjaneclark/buddy-api/app/models"
	"math/big"
	"os"
)
var DB_CONFIG = map[string]string{
	"host" :os.Getenv("DB_HOSTNAME"),			//The hostname of the database to connect to
	"port" : "5432",								//The port to connect to the db on. 5432 by default on postgres
	"user" : os.Getenv("DB_USER"),				//The db user
	"password" : os.Getenv("DB_PASSWORD"),		//The password to connect with
	"dbname" : os.Getenv("DB_NAME"),			//The name of the db to connect to
}

func dbConnect()*sql.DB{
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+ "password=%s dbname=%s",
		DB_CONFIG["host"], DB_CONFIG["port"], DB_CONFIG["user"], DB_CONFIG["password"], DB_CONFIG["dbname"])

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	return db
}

/*
	createUser
	This is the function that creates the user.
	It first connects to the database, then gets the data given by the session's user object,
	and inserts those into the database as a new user.
*/
func createUser(userData map[string]interface{}){
	db := dbConnect()

	if val, ok := userData["given_name"]; !ok {
		userData["first_name"] = ""
		fmt.Println(val)
	}
	if val, ok := userData["family_name"]; !ok {
		userData["family_name"] = ""
		fmt.Println(val)
	}
	if val, ok := userData["sub"]; !ok {
		userData["sub"] = ""
		fmt.Println(val)
	}

	if len(userData) != 0{
		sqlStatement := `
			INSERT INTO users (id, firstname, lastname )
			VALUES ('`+ userData["sub"].(string) +`',
					'`+ userData["given_name"].(string)+`',
					'`+ userData["family_name"].(string)+`');`
	db.Exec(sqlStatement)

	}
	defer db.Close()
}
/*
	getUserData
	Gets the userData from a returning user.
	It connects to the db and gets the user that exists with the id.
*/
func getUserData(userData map[string]interface{}, user *models.User) *models.User {
	db:= dbConnect()

	sqlStatement := `SELECT firstname, lastname FROM users WHERE id=$1;`
	row:= db.QueryRow(sqlStatement, userData["sub"])
	id := new(big.Int)
	id, ok := id.SetString(userData["sub"].(string), 10)
	if !ok {
		fmt.Println("SetString: error")
		return nil
	}
	user.Id = *id

	switch err := row.Scan(&user.Firstname, &user.Lastname); err{
		case sql.ErrNoRows:
			fmt.Print("")
		case nil:
			fmt.Print( user.Firstname, user.Lastname)
		default:
			panic(err)
	}
	defer db.Close()

	return user
}

func StoreInstanceId (user models.User, instanceId string, origin string){
	db := dbConnect()
	sqlStatement := `
			UPDATE InstanceIds SET (instanceId, origin) = (`+
			instanceId + `,` + origin +`) WHERE id = `+ user.Id.String() +`;`
	db.Exec(sqlStatement)
	defer db.Close()
}