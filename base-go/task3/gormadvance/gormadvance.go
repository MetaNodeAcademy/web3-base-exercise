package gormadvance

import (
	"log"
	"time"

	"github.com/LX/base-go/internal/db"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint      `gorm:"primaryKey"`
	Username string    `gorm:"column:username"`
	Password string    `gorm:"column:password"`
	Email    string    `gorm:"column:email"`
	Age      uint8     `gorm:"column:age"`
	Birthday time.Time `gorm:"column:birthday"`
	PostSize uint      `gorm:"column:post_size"`
	Posts    []Post    `gorm:"foreignKey:UserID"` // 一个用户有多篇文章
}

type Post struct {
	gorm.Model
	Title         string    `gorm:"column:title"`
	Content       string    `gorm:"column:content"`
	UserID        uint      `gorm:"column:user_id"` // 文章作者
	Comments      []Comment `gorm:"foreignKey:PostID;"`
	CommentStatus string    `gorm:"column:comment_status"`
}

type Comment struct {
	gorm.Model
	ID      uint   `gorm:"primaryKey"`
	Content string `gorm:"column:content"`
	UserID  uint   `gorm:"column:user_id"` // 评论作者
	PostID  uint   `gorm:"column:post_id"` // 评论所属文章
}

func GormAdvanceExercise() {
	log.Println("--------GormAdvanceExercise---------")
	db := db.InitDB()
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalln(err)
	}
	defer sqlDB.Close()
	/*
		题目1：模型定义
		假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
		要求 ：
		使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
		编写Go代码，使用Gorm创建这些模型对应的数据库表。
	*/
	//db.AutoMigrate(&User{}, &Post{}, &Comment{})
	// user := User{Username: "张三", Age: 18}
	// post1 := Post{Title: "文章1", Content: "这是文章1的内容"}
	// post2 := Post{Title: "文章2", Content: "这是文章2的内容"}
	// comment1 := Comment{Content: "评论1"}
	// comment2 := Comment{Content: "评论2"}
	// comment3 := Comment{Content: "评论3"}

	// post1.Comments = append(post1.Comments, comment1, comment2)
	// post2.Comments = append(post2.Comments, comment3)
	// user.Posts = append(user.Posts, post1, post2)
	// user.Birthday = time.Now()
	// result := db.Debug().Create(&user)
	// if result.Error != nil {
	// 	log.Fatalf("创建用户失败：%v", result.Error)
	// }

	// 题目2
	//user := GetPostByUserId(db, 1)
	//log.Println(user)
	//post := GetPostWithMostComments(db)
	//log.Println(post)

	// 题目3
	HookExercise(db)
}

/*
题目2：关联查询
基于上述博客系统的模型定义。
要求 ：
编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
编写Go代码，使用Gorm查询评论数量最多的文章信息。
*/
func GetPostByUserId(db *gorm.DB, userId uint) User {
	var user User
	db.Debug().
		Model(&User{ID: userId}).
		Preload("Posts").
		Preload("Posts.Comments").
		First(&user)
	return user
}

// 编写Go代码，使用Gorm查询评论数量最多的文章信息。
func GetPostWithMostComments(db *gorm.DB) Post {
	var post Post
	most := db.Debug().
		Model(&Comment{}).
		Select("post_id, COUNT( post_id ) AS comment_num").
		Group("post_id").
		Order("comment_num DESC").
		Limit(1)
	//关联子查询，查询出文章信息
	db.Debug().
		Model(&Post{}).
		Select("posts.*").
		Joins("JOIN (?) most ON posts.id = most.post_id", most).
		First(&post)
	db.Debug().Model(&Comment{PostID: post.ID}).Find(&post.Comments)
	return post
}

/*
题目3：钩子函数
继续使用博客系统的模型。
要求 ：
为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
*/
func HookExercise(db *gorm.DB) {
	log.Println("--------钩子函数-------")
	// 创建文章
	// user := User{Username: "张三", Age: 18}
	// post1 := Post{Title: "文章1", Content: "这是文章1的内容"}
	// post2 := Post{Title: "文章2", Content: "这是文章2的内容"}
	// comment1 := Comment{Content: "评论1"}
	// comment2 := Comment{Content: "评论2"}
	// comment3 := Comment{Content: "评论3"}

	// post1.Comments = append(post1.Comments, comment1, comment2)
	// post2.Comments = append(post2.Comments, comment3)
	// user.Posts = append(user.Posts, post1, post2)
	// user.Birthday = time.Now()
	// result := db.Debug().Create(&user)
	// if result.Error != nil {
	// 	log.Fatalf("创建用户失败：%v", result.Error)
	// }
	// log.Println(user)

	// 删除评论
	var comment Comment
	db.First(&comment, 51)
	db.Delete(&comment)
}

// 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
func (post *Post) AfterCreate(db *gorm.DB) (err error) {
	log.Println("--------AfterCreate-------")
	// 统计数量
	var count int64
	if error := db.Model(&Post{}).Where("user_id = ?", post.UserID).Count(&count).Error; error != nil {
		return error
	}
	// 更新回去
	if error := db.Model(&User{}).Where("id = ?", post.UserID).Update("post_size", uint(count)).Error; error != nil {
		return error
	}
	return nil
}

// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
func (comment *Comment) AfterDelete(db *gorm.DB) (err error) {
	log.Println("--------AfterDelete-------")
	// 统计数量
	var count int64
	if error := db.Model(&Comment{}).Where("post_id = ?", comment.PostID).Count(&count).Error; error != nil {
		return error
	}
	// 0 则更新
	if count == 0 {
		if error := db.Model(&Post{}).Where("id = ?", comment.PostID).Update("comment_status", "无评论").Error; error != nil {
			return error
		}
	}
	return nil
}