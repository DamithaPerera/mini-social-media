package routes

import (
	"mini-social-media/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterPostRoutes(router *gin.Engine) {
	postGroup := router.Group("/posts")
	{
		postGroup.POST("/", controllers.CreatePost)
		postGroup.PUT("/:id", controllers.UpdatePost)
		postGroup.GET("/", controllers.GetAllPosts)
		postGroup.POST("/:id/like", controllers.LikePost)
		postGroup.POST("/:id/comment", controllers.AddComment)
		postGroup.GET("/:id", controllers.GetPostDetails)
	}
}
