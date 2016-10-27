package db

import (
	"database/sql"
	"log"
	"fmt"
	"github.com/lib/pq"
	"github.com/sgjp/musico/util"
	"strconv"
	"time"
	"strings"
)

func GetUser(username, password string) int {

	db := getConnection()
	defer db.Close()
	var id int
	err := db.QueryRow("SELECT id FROM users WHERE username = $1 and password=$2", username, password).Scan(&id)

	switch {
	case err == sql.ErrNoRows:
		return -1
	case err != nil:
		log.Fatal(err)
		return -1
	default:
		user := User{id, username, password}
		fmt.Printf("User %v is logged in\n", user.Username)
		return id
	}
}

func GetBandByName(bandName string) Band {
	db := getConnection()
	defer db.Close()

	lowerName := strings.ToLower(bandName)

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
	bookings := make([]Booking, 0)

	err := db.QueryRow("SELECT * FROM bands WHERE lower(name)=$1;", lowerName).Scan(&id, &name, &genre, &youtube, &facebook, &requirements, &location, &avgPrice)
	util.CheckErr(err)

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

	rowsC, err := db.Query("SELECT id, comment, type FROM comments WHERE band_id=$1;", id)
	util.CheckErr(err)

	//Get the comments
	for rowsC.Next() {
		var id string
		var comment string
		var cType string
		err = rowsC.Scan(&id, &comment, &cType)
		cTypeI, _ := strconv.Atoi(cType)
		comm := Comment{id, comment, cTypeI}
		comments = append(comments, comm)
	}

	avgPriceI, _ := strconv.Atoi(avgPrice)

	rowsB, err := db.Query("SELECT id, description, date FROM bookings WHERE band_id=$1", id)
	util.CheckErr(err)

	//Get the bookings
	for rowsB.Next() {
		var id string
		var description string
		var date time.Time
		err = rowsB.Scan(&id, &description, &date)
		layout := "2006-01-02"
		dateT, _ := time.Parse(layout, date.Format(layout))
		booking := Booking{id, description, dateT}
		bookings = append(bookings, booking)
	}

	band := Band{id, name, genre, youtube, facebook, requirements, location, avgPriceI, reviews, comments, bookings, 5.0}
	band.AvgRate = band.GetAvgRate()

	return band
}



func AddBand(name, genre, avgPrice, location, youtube, facebook, requirements string) int {
	db := getConnection()
	defer db.Close()

	var id int

	err := db.QueryRow("INSERT INTO bands(name, genre, youtube, facebook, technical_reqs, location, avg_price) VALUES($1,$2,$3,$4,$5,$6,$7) returning id;",
		name, genre, youtube, facebook, requirements, location, avgPrice).Scan(&id)
	util.CheckErr(err)
	return id
}

func AddReview(comment string, rateQuality, ratePunctuality, rateFlexibility, rateEnthusiasm, rateSimilarity, rate, userId, bandId int) int {

	db := getConnection()
	defer db.Close()
	var id int

	err := db.QueryRow("INSERT INTO reviews(user_id, band_id, comment, rate_quality, rate_punctuality, rate_flexibility, rate_enthusiasm, rate_similarity, rate)" +
		"VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) returning id;", userId, bandId, comment, rateQuality, ratePunctuality, rateFlexibility, rateEnthusiasm, rateSimilarity, rate).Scan(&id)
	util.CheckErr(err)
	return id

}

func AddComment(comment string, cType, userId, bandId int) int {

	db := getConnection()
	defer db.Close()
	var id int

	err := db.QueryRow("INSERT INTO comments(band_id, user_id, type, comment) VALUES ($1,$2,$3,$4) returning id;", bandId, userId, cType, comment).Scan(&id)
	util.CheckErr(err)
	return id

}

func GetAllBands() []Band {
	bands := make([]Band, 0)
	db := getConnection()
	defer db.Close()
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
		bookings := make([]Booking, 0)

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

		rowsC, err := db.Query("SELECT id, comment, type FROM comments WHERE band_id=$1;", id)

		//Get the comments
		for rowsC.Next() {
			var id string
			var comment string
			var cType string
			err = rowsC.Scan(&id, &comment, &cType)
			cTypeI, _ := strconv.Atoi(cType)
			comm := Comment{id, comment, cTypeI}
			comments = append(comments, comm)
		}

		avgPriceI, _ := strconv.Atoi(avgPrice)

		rowsB, err := db.Query("SELECT id, description, date FROM bookings WHERE band_id=$1", id)

		//Get the bookings
		for rowsB.Next() {
			var id string
			var description string
			var date time.Time
			err = rowsB.Scan(&id, &description, &date)
			layout := "2006-01-02"
			dateT, _ := time.Parse(layout, date.Format(layout))
			booking := Booking{id, description, dateT}
			bookings = append(bookings, booking)
		}

		band := Band{id, name, genre, youtube, facebook, requirements, location, avgPriceI, reviews, comments, bookings, 5.0}
		band.AvgRate = band.GetAvgRate()
		bands = append(bands, band)

	}

	return bands

}

func AddBooking(description, date, bandIdS string) int {

	bandId, err := strconv.Atoi(bandIdS)
	util.CheckErr(err)

	layout := "2006-01-02"
	dateD, _ := time.Parse(layout, date)

	db := getConnection()
	defer db.Close()
	var id int
	err = db.QueryRow("INSERT INTO bookings(description, date, band_id) VALUES ($1, $2, $3) returning id;", description, dateD, bandId).Scan(&id)
	util.CheckErr(err)

	return id
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
