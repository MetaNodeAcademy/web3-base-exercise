package blog

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/LX/base-go/internal/db"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var gdb *gorm.DB
var jwtSecret = []byte("mDdL7RnUPK4l/qOiWtKlyNLkvg9X+wpfVft3jmfOtmo=")

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"not null"`
}

type Post struct {
	gorm.Model
	Title   string `gorm:"not null"`
	Content string `gorm:"not null"`
	UserID  uint
	User    User `gorm:"foreignkey:UserID"`
}

type Comment struct {
	gorm.Model
	Content string `gorm:"not null"`
	UserID  uint
	User    User `gorm:"foreignkey:UserID"`
	PostID  uint `json:"postID"`
	Post    Post `gorm:"foreignkey:PostID"`
}

// 模拟主线程
func BlogExercise() {
	log.Println("--------Blog Exercise--------")
	gdb = db.InitDB()
	sqlDB, err := gdb.DB()
	if err != nil {
		log.Fatalln(err)
	}
	defer sqlDB.Close()
	gdb.AutoMigrate(&User{}, &Post{}, &Comment{})
	// 启动路由
	r := InitRouter()
	r.Run(":8080")
}

// 注册
func Register(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = string(hashedPassword)
	if err := gdb.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

// 登陆
func Login(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var storedUser User
	if err := gdb.Model(&User{}).Where("username = ?", user.Username).First(&storedUser).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}
	// 验证密码 第一个参数: 数据库里的密码哈希值 第二个参数: 用户输入的密码
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}
	// 生成 JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       storedUser.ID,
		"username": storedUser.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	log.Println("Generated token:", tokenString)
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

// 验证Token中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头获取 Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}
		// 检查是否符合 Bearer 格式
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer {token}"})
			c.Abort()
			return
		}
		tokenString := parts[1]
		// 解析 JWT
		token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
			// 这里返回的是你的密钥，用于校验签名
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}
		// 从 claims 里取数据，存进 Gin context
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			// c.Set("userID", claims["id"])
			c.Set("userID", uint(claims["id"].(float64)))
			c.Set("username", claims["username"])
		}
		c.Next()
	}
}
