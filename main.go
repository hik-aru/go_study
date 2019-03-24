package main

import (
	"dev/controllers"
	"github.com/gin-gonic/gin"
	"reflect"
	"strconv"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/index/:id", func(c *gin.Context) {
		//parmを処理する
		n := c.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil{
			c.JSON(400, err)
			return
		}
		if id <= 0 {
			c.JSON(400, gin.H{"status": "id should be bigger than 0"})

			return
		}
		//データを処理する
		ctrl := controllers.NewUser()
		result := ctrl.Get(id)
		if result == nil || reflect.ValueOf(result).IsNil() {
			c.JSON(404, gin.H{})
			return
		}
		
		c.JSON(200, result)
	})

	router.LoadHTMLGlob("templates/*")
	router.GET("/test", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	router.Run(":8080")
}
