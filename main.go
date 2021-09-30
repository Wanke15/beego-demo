package main

import (
	"github.com/yanyiwu/gojieba"

	"github.com/gin-gonic/gin"
)

func main() {

	use_hmm := true
	x := gojieba.NewJieba()
	defer x.Free()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/cut", func(c *gin.Context) {
		sent := c.DefaultQuery("sent", "")
		var words []string = x.Cut(sent, use_hmm)

		c.JSON(200, gin.H{
			"words": words,
		})
	})

	r.Run("127.0.0.1:8089") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
