package blog

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type QueryParams struct {
	PostID int `form:"id" binding:"required"`
}

// 创建文章 实现文章的创建功能，只有已认证的用户才能创建文章，创建文章时需要提供文章的标题和内容
func CreatePost(c *gin.Context) {
	userId := c.GetUint("userID")
	userName, _ := c.Get("username")
	post := Post{}
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	addPost := Post{
		Title:   post.Title,
		Content: post.Content,
		User: User{
			Model: gorm.Model{
				ID: userId,
			},
			Username: userName.(string),
		},
	}
	gdb.Create(&addPost)
	c.JSON(200, gin.H{"message": "添加成功", "data": post.Content})
	log.Println("addPost %+v", addPost)
}

// 读取文章 实现文章的读取功能，支持获取所有文章列表和单个文章的详细信息
func GetPost(c *gin.Context) {
	postIdString := c.Query("id")
	postId, err := strconv.Atoi(postIdString)
	if err != nil {
		fmt.Println("转换失败:", err)
		// 给默认值后获取全部文章
		postId = 0
	}
	// 如果id为空, 则获取所有文章
	if postId == 0 {
		var posts []Post
		gdb.Find(&posts)
		c.JSON(200, gin.H{"message": "获取成功", "data": posts})
		return
	} else {
		var post Post
		gdb.Where("id = ?", postId).First(&post)
		c.JSON(200, gin.H{"message": "获取成功", "data": post})
		return
	}
}

// 更新文章 实现文章的更新功能，只有文章的作者才能更新自己的文章
func UpdatePost(c *gin.Context) {
	userId := c.GetUint("userID")
	postJson := Post{}
	post := Post{}
	if err := c.ShouldBindJSON(&postJson); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := gdb.Where("id = ?", postJson.ID).First(&post).Error; err != nil {
		c.JSON(404, gin.H{"error": "文章不存在"})
		return
	}
	if post.UserID != userId {
		c.JSON(403, gin.H{"error": "无权限更新文章"})
		return
	}
	if err := gdb.Model(&postJson).Where("id = ?", postJson.ID).Updates(postJson).Error; err != nil {
		c.JSON(500, gin.H{"error": "更新文章失败"})
		return
	}
	c.JSON(200, gin.H{"message": "文章更新成功"})
}

// 删除文章 实现文章的删除功能，只有文章的作者才能删除自己的文章
func DeletePost(c *gin.Context) {
	userId := c.GetUint("userID")
	var params QueryParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid query params"})
		return
	}
	postId := params.PostID
	if err := gdb.Where("id = ? and user_id = ?", postId, userId).Delete(&Post{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "delete post failed"})
		return
	}
	log.Println("userId: %d, postId: %s", userId, postId)
}
