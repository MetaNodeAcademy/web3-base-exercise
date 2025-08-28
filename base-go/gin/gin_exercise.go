package gin

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	User     string `json:"user" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func GinExercise() {
	log.Println("--------Gin Exercise Started---------")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("login", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "login",
		})
	})
	// 路由组 v1
	{
		v1 := r.Group("/v1")
		v1.POST("login", login(v1.BasePath()))
		v1.POST("submit", submit(v1.BasePath()))
		v1.POST("read", read(v1.BasePath()))
	}
	// 路由组 v2
	{
		v2 := r.Group("/v2")
		v2.POST("login", login(v2.BasePath()))
		v2.POST("submit", submit(v2.BasePath()))
		v2.POST("read", read(v2.BasePath()))
	}
	// r.GET("/user", login("/user"))    //查询
	// r.POST("/user", login("/user"))   // 新增
	// r.DELETE("/user", login("/user")) // 删除
	// r.PUT("/user", login("/user"))    // 更新（客户端提供完整数据）
	// r.PATCH("/user", login("/user"))  // 更新（客户端提供需要修改的数据）
	// 重定向到外部
	// r.GET("/test1", func(c *gin.Context) {
	// 	c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com/")
	// })

	// 重定向到内部
	// r.POST("/test4", func(c *gin.Context) {
	// 	c.Redirect(http.StatusFound, "/login")
	// })
	// r.GET("/test5", func(c *gin.Context) {
	// 	c.Request.URL.Path = "/ping"
	// 	r.HandleContext(c)
	// })
	// r.GET("/test6", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{"hello": "world"})
	// })
	// JSON参数
	r.POST("/loginJSON", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if json.User != "root" || json.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "登陆成功"})
	})
	// 静态文件服务
	//assetsPath, _ := filepath.Abs("./assets")
	//log.Println("路径: %s", assetsPath)
	// 访问 http://127.0.0.1:8080/static/doc.txt
	r.Static("/static", "../../assets")
	// 访问 http://127.0.0.1:8080/static2/doc.txt
	r.StaticFS("/static2", http.Dir("../../assets"))
	// 访问 http://127.0.0.1:8080/static3
	r.StaticFile("/static3", "../../assets/doc.txt")

	r.Run(":8080")
}

func login(s string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "login",
			"version": s,
		})
	}
}

func submit(s string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "submit v1",
			"version": s,
		})
	}
}

func read(s string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "read v1",
			"version": s,
		})
	}
}
