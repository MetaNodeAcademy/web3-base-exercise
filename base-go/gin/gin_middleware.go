package gin

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// 模拟一些私人数据
var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

func GinMiddleware() {
	log.Println("--------GinMiddleware start---------")
	// r := gin.New()
	// r.Use(Logger())
	// r.GET("/test", func(c *gin.Context) {
	// 	example := c.MustGet("example").(string)
	// 	// 打印
	// 	log.Println(example)
	// })
	// r.Run(":8080")

	// 课件
	// r := gin.Default()
	// r.GET("/", mw1(), mw2(), func(c *gin.Context) {
	// 	log.Println("self")
	// 	c.String(http.StatusOK, "self")
	// })
	// r.Run(":8080")

	// r := gin.Default()
	// // 路由组使用 gin.BasicAuth() 中间件
	// authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
	// 	"foo":    "bar",
	// 	"austin": "123",
	// 	"lena":   "hello2",
	// 	"manu":   "password123",
	// }))
	// authorized.GET("/secrets", func(c *gin.Context) {
	// 	// 获取用户，它是由 BasicAuth 中间件设置的
	// 	user := c.MustGet(gin.AuthUserKey).(string)
	// 	if secret, ok := secrets[user]; ok {
	// 		c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
	// 	} else {
	// 		c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
	// 	}
	// })
	// r.Run(":8080")

	// 新建一个没有任何默认中间件的路由
	//r := gin.New()
	// 全局中间件
	// Logger 中间件将日志写入 gin.DefaultWriter，即使你将 GIN_MODE 设置为 release。
	// By default gin.DefaultWriter = os.Stdout
	//r.Use(gin.Logger())
	// Recovery 中间件会 recover 任何 panic。如果有 panic 的话，会写入 500。
	//r.Use(gin.Recovery())
	// 你可以为每个路由添加任意数量的中间件。
	//r.GET("/benchmark", MyBenchLogger(), benchEndpoint)

	// 认证路由组
	// authorized := router.Group("/", AuthRequired())
	// 和使用以下两行代码的效果完全一样:
	//authorized := r.Group("/")
	// 路由组中间件! 在此例中，我们在 "authorized" 路由组中使用自定义创建的
	// AuthRequired() 中间件
	//authorized.Use(AuthRequired())
	// {
		// authorized.POST("/login", loginEndpoint)
		// authorized.POST("/submit", submitEndpoint)
		// authorized.POST("/read", readEndpoint)
		// 嵌套路由组
		// testing := authorized.Group("testing")
		// testing.GET("/analytics", analyticsEndpoint)
	// }
	// 监听并在 0.0.0.0:8080 上启动服务
	// r.Run(":8080")
	
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		// 设置 example 变量
		c.Set("example", "12345")
		// 请求前
		c.Next()
		// 请求后
		latency := time.Since(t)
		log.Print(latency)
		// 获取发送的 status
		status := c.Writer.Status()
		log.Println(status)
	}
}

func mw1() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("mw1 before")
		c.Next()
		log.Println("mw1 after")
	}
}

func mw2() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("mw2 before")
		c.Next()
		log.Println("mw2 after")
	}
}
