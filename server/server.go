package server

import (
	"github.com/gin-gonic/gin"
	//"net/http"
	"github.com/sgjp/musico/db"
	"strconv"
	"github.com/sgjp/musico/util"
	"strings"
	"log"
)

func StartServer() {
	router := gin.Default()

	router.GET("/test", test)

	router.POST("/auth/login", login)

	router.POST("/band", addBand)

	router.GET("/bands", getAllBands)

	router.GET("/bands/search", searchBands)

	router.POST("/band/:id/review", addReview)

	router.POST("/band/:id/comment", addComment)

	router.POST("/band/:id/booking", addBooking)

	router.GET("/info", func(c *gin.Context) {
		c.JSON(200, router.Routes())
	})

	router.Run()
}

func login(c *gin.Context) {
	userName := c.PostForm("userName")
	password := c.PostForm("password")

	result := db.GetUser(userName, password)
	if result {
		content := gin.H{"userName": userName}
		c.JSON(200, content)
	} else {
		c.Status(404)
	}

}

func addBand(c *gin.Context) {
	name := c.PostForm("name")
	genre := c.PostForm("genre")
	avgPrice := c.PostForm("avgPrice")
	requirements := c.PostForm("requirements")
	location := c.PostForm("location")
	youtube := c.PostForm("youtube")
	facebook := c.PostForm("facebook")

	id := db.AddBand(name, genre, avgPrice, location, youtube, facebook, requirements)

	if id > 0 {
		content := gin.H{"id": id}
		c.JSON(200, content)
	} else {
		c.Status(500)
	}

}

func getAllBands(c *gin.Context) {
	bands := db.GetAllBands()

	if len(bands) > 0 {
		c.JSON(200, bands)
	} else {
		c.Status(204)

	}

}

func searchBands(c *gin.Context) {
	minPrice := c.DefaultQuery("minPrice", "0")
	maxPrice := c.DefaultQuery("maxPrice", "999999999")
	location := c.Query("location")
	genre := c.Query("genre")
	minAvgRate := c.DefaultQuery("minRate", "0")
	availableDate := c.Query("availableDate")

	bands := db.GetAllBands()
	bandsFiltered := make([]db.Band, 0)

	for i := range bands {
		if (bands[i].AvgPrice >= util.ToInt(minPrice) && bands[i].AvgPrice <= util.ToInt(maxPrice)) || bands[i].AvgPrice == 0 {
			if location == "" || strings.ToLower(bands[i].Location) == strings.ToLower(location) {
				avgRate := bands[i].AvgRate
				minAvgRateF, err := strconv.ParseFloat(minAvgRate, 64)

				if (err != nil) {
					log.Printf("Error converting %v into float", minAvgRate)
					c.Status(500)
					return
				}
				if minAvgRateF == 0.0 || avgRate >= minAvgRateF {
					if bands[i].IsAvailable(availableDate) {
						if genre == "" || strings.ToLower(bands[i].Genre) == strings.ToLower(genre) {
							bandsFiltered = append(bandsFiltered, bands[i])
						}
					}
				}
			}

		}

	}

	if len(bandsFiltered) > 0 {
		c.JSON(200, bandsFiltered)
	} else {
		c.Status(204)

	}

}

func addReview(c *gin.Context) {
	bandId, err := strconv.Atoi(c.Param("id"))
	util.CheckErr(err)

	comment := c.PostForm("comment")
	rateQuality, err := strconv.Atoi(c.PostForm("rateQuality"))
	util.CheckErr(err)
	ratePunctuality, err := strconv.Atoi(c.PostForm("ratePunctuality"))
	util.CheckErr(err)
	rateFlexibility, err := strconv.Atoi(c.PostForm("rateFlexibility"))
	util.CheckErr(err)
	rateEnthusiasm, err := strconv.Atoi(c.PostForm("rateEnthusiasm"))
	util.CheckErr(err)
	rateSimilarity, err := strconv.Atoi(c.PostForm("rateSimilarity"))
	util.CheckErr(err)
	rate, err := strconv.Atoi(c.PostForm("rate"))
	util.CheckErr(err)
	userId, err := strconv.Atoi(c.PostForm("userId"))
	util.CheckErr(err)

	id := db.AddReview(comment, rateQuality, ratePunctuality, rateFlexibility, rateEnthusiasm, rateSimilarity, rate, userId, bandId)

	if id > 0 {
		content := gin.H{"id": id}
		c.JSON(200, content)
	} else {
		c.Status(500)
	}
}

func addComment(c *gin.Context) {
	bandId, err := strconv.Atoi(c.Param("id"))
	util.CheckErr(err)

	comment := c.PostForm("comment")
	cType, err := strconv.Atoi(c.PostForm("type"))
	util.CheckErr(err)
	userId, err := strconv.Atoi(c.PostForm("userId"))
	util.CheckErr(err)

	id := db.AddComment(comment, cType, userId, bandId)

	if id > 0 {
		content := gin.H{"id": id}
		c.JSON(200, content)
	} else {
		c.Status(500)
	}
}

func addBooking(c *gin.Context) {
	bandId := c.Param("id")

	description := c.PostForm("description")
	date := c.PostForm("date")

	id := db.AddBooking(description, date, bandId)

	if id > 0 {
		content := gin.H{"id": id}
		c.JSON(200, content)
	} else {
		c.Status(500)
	}
}

func test(c *gin.Context) {
	content := gin.H{"Status":"OK"}
	c.JSON(200, content)
}