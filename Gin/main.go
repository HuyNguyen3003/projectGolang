package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"

	"github.com/HuyNguyen/gin/initializer"
)

func initEnv() {

	initializer.LoadEnvVarible()
	initializer.InitMongodb()

}

func main() {
	initEnv()
	r := gin.Default()

	r.POST("/user", func(c *gin.Context) {
		u := initializer.User{}
		if err := c.Bind(&u); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		mgm.Coll(&u).Create(&u)
		//mgm.Coll(u).Create(u)
		c.JSON(http.StatusOK, gin.H{
			"user": u,
		})
	})
	r.Run("127.0.0.1:5000")
}
