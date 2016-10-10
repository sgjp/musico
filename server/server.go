package server

import (
	"github.com/gin-gonic/gin"
	//"net/http"
	"github.com/sgjp/musico/db"
	"strconv"
	"github.com/sgjp/musico/util"
)

func StartServer() {
	router := gin.Default()

	router.GET("/test", test)

	router.POST("/auth/login", login)

	router.POST("/band", addBand)

	router.GET("/bands", getAllBands)

	router.POST("/band/:id/review", addReview)

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

func addReview(c *gin.Context) {
	bandId,err := strconv.Atoi(c.Param("id"))
	util.CheckErr(err)

	comment := c.PostForm("comment")
	rateQuality,err := strconv.Atoi(c.PostForm("rateQuality"))
	util.CheckErr(err)
	ratePunctuality,err := strconv.Atoi(c.PostForm("ratePunctuality"))
	util.CheckErr(err)
	rateFlexibility,err := strconv.Atoi(c.PostForm("rateFlexibility"))
	util.CheckErr(err)
	rateEnthusiasm,err := strconv.Atoi(c.PostForm("rateEnthusiasm"))
	util.CheckErr(err)
	rateSimilarity,err := strconv.Atoi(c.PostForm("rateSimilarity"))
	util.CheckErr(err)
	rate,err := strconv.Atoi(c.PostForm("rate"))
	util.CheckErr(err)
	userId,err := strconv.Atoi(c.PostForm("userId"))
	util.CheckErr(err)

	id := db.AddReview(comment,rateQuality, ratePunctuality, rateFlexibility, rateEnthusiasm, rateSimilarity, rate, userId, bandId)

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