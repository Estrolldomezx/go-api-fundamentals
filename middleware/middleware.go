package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("middleware", exampleMiddleware, exampleHandler)

	// route
	group1 := r.Group("/group1")
	group1.GET("/handler1", handler1)
	group1.GET("/handler1", handler1)

	group2 := r.Group("/group2")
	group2.Use(exampleMiddleware)
	group2.GET("/handler3", handler3)
	r.Run()
}

func exampleMiddleware(c *gin.Context) {
	fmt.Println("start middleware")
	c.Next()
	fmt.Println("end middleware")
}

func exampleHandler(c *gin.Context) {
	fmt.Println("this is handler")
}

func handler1(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "this is handler1",
	})
}

func handler2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "this is handler2",
	})
}

func handler3(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "this is handler3",
	})
}
