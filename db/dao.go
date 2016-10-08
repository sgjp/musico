package db

import (
	"database/sql"
	"log"
	"fmt"
	"github.com/lib/pq"
)

func GetUser(username, password string) bool {

	db := getConnection()
	var id int
	var userName string
	var pass string
	err := db.QueryRow("SELECT * FROM users WHERE username = $1 and password=$2", username, password).Scan(&id,&username,&pass)

	switch {
	case err == sql.ErrNoRows:
		return false
	case err != nil:
		log.Fatal(err)
		return false
	default:
		user := User{id, userName, pass}
		fmt.Printf("User is %v logged in\n", user.Username)
		return true
	}
}

func GetReviews() {
	/*for rows.Next() {
		var uid int
		var username string
		var department string
		var created time.Time
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println("uid | username | department | created ")
		fmt.Printf("%3v | %8v | %6v | %6v\n", uid, username, department, created)
	}*/
}

func getConnection() *sql.DB {
	/*url, _ := pq.ParseURL("ec2-23-23-225-158.compute-1.amazonaws.com")
	db, err := sql.Open(
		"postgres", "host=" + url +
			"user=ruhohuczvwkoho " +
			"dbname=d3pnlqjqplr3q4 " +
			"password=_owsxMjfR3gtMELsO8Og4EldB6" +
			"sslmode=require")*/
	conn, _ := pq.ParseURL("postgres://ruhohuczvwkoho:_owsxMjfR3gtMELsO8Og4EldB6@ec2-23-23-225-158.compute-1.amazonaws.com/d3pnlqjqplr3q4?sslmode=require")
	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return db
}
