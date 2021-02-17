package main

import (
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/", queryByString)
	http.ListenAndServe(":9000", r)
}

func queryByString(c *gin.Context) {
	q := c.Query("q")
	from := c.DefaultQuery("from", "0")
	pageFrom, _ := strconv.Atoi(from)
	result, err := QueryData(q, pageFrom)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "query error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "OK",
		"data": result,
	})
}
