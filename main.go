package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"fmt"
)

func main() {
	r := gin.Default()
	r.GET("/pcc", func(c *gin.Context) {
		action := c.Param("action")
		switch action {
		case "like":
			like(c)
		case "is_like":
			isLike(c)
		case "count":
			count(c)
		case "list":
			list(c)
		default:
			c.JSON(200, gin.H{
				"error_code": 400,
				"error_message":fmt.Sprintf("unexpected action '%s'", action),
			})
		}

	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
