package request

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Customer struct {
	id   int
	name string
	male string
}

var customer = []Customer{
	{
		id:   1,
		name: "Dome",
		male: "male",
	}, {
		id:   2,
		name: "Dook",
		male: "female",
	}, {
		id:   3,
		name: "Deek",
		male: "female",
	},
}

func Request() {
	r := gin.Default()
	r.GET("/", Hello)
	r.GET("/path/:name", pathParam)
	r.GET("/query", queryParam)
	r.GET("/body", bodyParam)
	r.Run()
}

// handler function
func Hello(c *gin.Context) {
	message := "success"
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}

// only in file
func pathParam(c *gin.Context) {
	param := c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"param": param,
	})
}

// query params
func queryParam(c *gin.Context) {
	param1 := c.Query("p1")
	param2 := c.Query("p2")
	c.JSON(http.StatusOK, gin.H{
		"param1": param1,
		"param2": param2,
	})
}

func bodyParam(c *gin.Context) {
	var request Customer

	err := c.ShouldBind(&request)
	// nil == null
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "bad request",
		})
		return
	}
	c.JSON(http.StatusOK, request)
}
