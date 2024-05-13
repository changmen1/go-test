package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		//指定默认值
		//http://localhost:8080/user 才会打印出来默认的值
		name := c.DefaultQuery("name", "枯藤")
		c.String(http.StatusOK, fmt.Sprintf("hello %s", name))
	})
	r.GET("/name", func(c *gin.Context) {
		c.String(http.StatusOK, fmt.Sprintf("李瑶"))
	})
	err := r.Run(":8060")
	if err != nil {
		return
	}
}
