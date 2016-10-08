package server

import (
	"github.com/gin-gonic/gin"
	//"net/http"
	"github.com/sgjp/musico/db"
)

func StartServer(){
	router := gin.Default()

	router.POST("/auth/login", login)

	router.Run()
}

func login (c *gin.Context){
	userName := c.PostForm("userName")
	password := c.PostForm("password")

	result := db.GetUser(userName,password)
	if result {
		content := gin.H{"userName": userName}
		c.JSON(200, content)
	}else{
		c.Status(404)
	}


}
