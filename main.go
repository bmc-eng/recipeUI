package main

import "github.com/gin-gonic/gin"

func IndexHandler(c *gin.Context) {
	c.File("index.html")
}

func main() {
	router := gin.Default()
	router.GET("/", IndexHandler)
	router.Run()
}
