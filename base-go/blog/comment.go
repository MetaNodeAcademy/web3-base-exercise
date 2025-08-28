package blog

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 实现评论的创建功能，已认证的用户可以对文章发表评论
func CreateComment(c *gin.Context) {
	userId := c.GetUint("userID")
	comment := Comment{}
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := gdb.Create(&Comment{
		Content: comment.Content,
		Post: Post{
			Model: gorm.Model{
				ID: comment.PostID,
			},
		},
		User: User{
			Model: gorm.Model{
				ID: userId,
			},
		},
	}).Error; err != nil {
		c.JSON(500, gin.H{"error": "创建评论失败"})
		return
	}
	c.JSON(200, gin.H{"message": "评论创建成功"})
}

// 实现评论的读取功能，支持获取某篇文章的所有评论列表
func GetComment(c *gin.Context) {
	var comments []Comment
	postIDString := c.Query("postID")
	postID, err := strconv.Atoi(postIDString)
	if err != nil {
		log.Println("转换失败:", err)
		// 给默认值后获取全部文章
		postID = 0
	}
	if err := gdb.Where("post_id = ?", postID).Find(&comments).Error; err != nil {
		c.JSON(500, gin.H{"error": "获取评论列表失败"})
		return
	}
	c.JSON(200, comments)
}
