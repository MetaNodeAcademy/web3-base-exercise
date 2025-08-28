package blog

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	r := gin.Default()
	apiGroup := r.Group("/api")
	{
		apiGroup.POST("/register", Register)
		apiGroup.POST("/login", Login)
	}
	// 需要授权的路由组
	apiAuthGroup := r.Group("/api", AuthMiddleware())
	{
		apiAuthGroup.POST("/post", CreatePost)
		apiAuthGroup.GET("/getPost", GetPost)
		apiAuthGroup.POST("/updatePost", UpdatePost)
		apiAuthGroup.DELETE("/deletePost", DeletePost)
		apiAuthGroup.POST("/createComment", CreateComment)
		apiAuthGroup.GET("/getComment", GetComment)
	}
	return r	
}