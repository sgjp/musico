package db

import (
	"time"
	"strconv"
)

type User struct {
	Id       int
	Username string
	Password string
}

type Band struct {
	Id           int
	Name         string
	Genre        string
	Youtube      string
	Facebook     string
	Requirements string
	Location     string
	AvgPrice     int
	Reviews      []Review
	Comments     []Comment
	Bookings     []Booking
	AvgRate	     float64
}

func (b Band) GetAvgRate() float64{
	if len(b.Reviews)==0{
		//Bands without reviews will get the maximum rating
		return 5
	}
	var avgRate int
	for i := range b.Reviews{
		rate,_ := strconv.Atoi(b.Reviews[i].Rate)
		avgRate += rate
	}
	return float64(avgRate)/float64(len(b.Reviews))
}

func (b Band) IsAvailable(date string) bool{
	if date==""{
		return true
	}

	layout := "2006-01-02"


	dateD, _ := time.Parse(layout, date)

	for i := range b.Bookings{
		bookedDate := b.Bookings[i].Date
		if bookedDate.Equal(dateD){
			return false
		}
	}
	return true
}


type Review struct {
	Id              string
	Comment         string
	RateQuality     string
	RatePunctuality string
	RateFlexibility string
	RateEnthusiasm  string
	RateSimilarity  string
	Rate            string
}

type Comment struct {
	Id      string
	Comment string
	Type    int
}

type Booking struct {
	Id          string
	Description string
	Date        time.Time
}