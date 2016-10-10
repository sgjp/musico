package db

import (
	"database/sql"
	"log"
	"fmt"
	"github.com/lib/pq"
	"github.com/sgjp/musico/util"
	"strconv"
)

func GetUser(username, password string) bool {

	db := getConnection()
	var id int
	var userName string
	var pass string
	err := db.QueryRow("SELECT * FROM users WHERE username = $1 and password=$2", username, password).Scan(&id, &username, &pass)

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

func AddBand(name, genre, avgPrice, location, youtube, facebook, requirements string) int {
	db := getConnection()

	var id int

	err := db.QueryRow("INSERT INTO bands(name, genre, youtube, facebook, technical_reqs, location, avg_price) VALUES($1,$2,$3,$4,$5,$6,$7) returning id;",
		name, genre, youtube, facebook, requirements, location, avgPrice).Scan(&id)
	util.CheckErr(err)
	return id
}

func AddReview(comment string, rateQuality, ratePunctuality, rateFlexibility, rateEnthusiasm, rateSimilarity, rate, userId, bandId int) int {

	db := getConnection()

	var id int

	err := db.QueryRow("INSERT INTO reviews(user_id, band_id, comment, rate_quality, rate_punctuality, rate_flexibility, rate_enthusiasm, rate_similarity, rate)" +
		"VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) returning id;", userId, bandId, comment, rateQuality, ratePunctuality, rateFlexibility, rateEnthusiasm, rateSimilarity, rate).Scan(&id)
	util.CheckErr(err)
	return id

}

func AddComment(comment string, cType, userId, bandId int) int {

	db := getConnection()

	var id int

	err := db.QueryRow("INSERT INTO comments(band_id, user_id, type, comment) VALUES ($1,$2,$3,$4) returning id;",  bandId, userId, cType, comment).Scan(&id)
	util.CheckErr(err)
	return id

}

func GetAllBands() []Band {
	bands := make([]Band, 0)
	db := getConnection()
	rows, err := db.Query("SELECT * FROM bands;")
	util.CheckErr(err)

	//Get all the bands
	for rows.Next() {
		var id int
		var name string
		var genre string
		var youtube string
		var facebook string
		var requirements string
		var location string
		var avgPrice string
		reviews := make([]Review, 0)
		comments := make([]Comment, 0)

		err = rows.Scan(&id, &name, &genre, &youtube, &facebook, &requirements, &location, &avgPrice)

		rowsR, err := db.Query("SELECT id, comment, rate_quality, rate_punctuality, rate_flexibility, rate_enthusiasm, rate_similarity, rate FROM reviews WHERE band_id=$1;", id)

		//Get the reviews
		for rowsR.Next() {
			var id string
			var comment string
			var rateQuality string
			var ratePunctuality string
			var rateFlexibility string
			var rateEnthusiasm string
			var rateSimilarity string
			var rate string
			err = rowsR.Scan(&id, &comment, &rateQuality, &ratePunctuality, &rateFlexibility, &rateEnthusiasm, &rateSimilarity, &rate)
			review := Review{id, comment, rateQuality, ratePunctuality, rateFlexibility, rateEnthusiasm, rateSimilarity, rate}
			reviews = append(reviews, review)
		}
		util.CheckErr(err)

		rowsC, err := db.Query("SELECT id, comment, type FROM comments WHERE band_id=$1;",id)

		for rowsC.Next(){
			var id string
			var comment string
			var cType string
			err = rowsC.Scan(&id, &comment, &cType)
			cTypeI,_ := strconv.Atoi(cType)
			comm := Comment{id,comment,cTypeI}
			comments = append(comments, comm)
		}

		avgPriceI,_:=strconv.Atoi(avgPrice)
		band := Band{id, name, genre, youtube, facebook, requirements, location, avgPriceI, reviews, comments}
		bands = append(bands, band)

	}

	return bands

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
	conn, _ := pq.ParseURL("postgres://ruhohuczvwkoho:_owsxMjfR3gtMELsO8Og4EldB6@ec2-23-23-225-158.compute-1.amazonaws.com/d3pnlqjqplr3q4?sslmode=require")
	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return db
}
