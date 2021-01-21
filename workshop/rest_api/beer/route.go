package beer

import "github.com/gin-gonic/gin"

func NewRoutes(r *gin.Engine) {
	b := r.Group("/beer")
	b.GET("/", getAllBeer)
}

func getAllBeer(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get all beer",
	})
}
