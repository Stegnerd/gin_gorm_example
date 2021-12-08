package api

import "github.com/gin-gonic/gin"

func GetRouter() *gin.Engine {
	r := gin.Default()

	return r
}

func AddPostRoutes(r *gin.Engine) {
	routes := r.Group("/posts")
	{
		getRoute(routes)
	}

}

func getRoute(r *gin.RouterGroup) {
	r.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
