package example1

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint           // Standard field for the primary key
	Name         string         // A regular string field
	Email        *string        // A pointer to a string, allowing for null values
	Age          uint8          // An unsigned 8-bit integer
	Birthday     *time.Time     // A pointer to time.Time, can be null
	MemberNumber sql.NullString // Uses sql.NullString to handle nullable strings
	ActivatedAt  sql.NullTime   // Uses sql.NullTime for nullable time fields
	CreatedAt    time.Time      // Automatically managed by GORM for creation time
	UpdatedAt    time.Time      // Automatically managed by GORM for update time
	ignored      string         // fields that aren't exported are ignored
}

type Member struct {
	gorm.Model
	Name string
	Age  uint8
}

type Author struct {
	Name  string
	Email string
}

type Blog struct {
	ID int
	Author
	Upvotes int32
}

type Blog2 struct {
	ID     int64
	Author Author `gorm:"embedded;embeddedPrefix:author_"`
	// Author Author
	Upvotes int32
}

func Run(db *gorm.DB) {
	fmt.Println("--------CRUD example1--------")
	// db.AutoMigrate(&User{})
	// db.AutoMigrate(&Member{})
	// db.AutoMigrate(&Blog{})
	// db.AutoMigrate(&Blog2{})

	// 创建
	// user := &User{}
	// user.Name = "ba"
	// user.Age = 28
	// user.MemberNumber.Valid = true
	// db.Create(user)

	// member := &Member{}
	// member.Name = "柒"
	// member.Age = 17
	// db.Create(&member)
	// fmt.Println("memberID = ", member.ID)
	//db.Delete(&Member{}, 2)
	//db.Delete(&User{}, 4)

	// ctx := context.Background()
	// user1, err := gorm.G[User](db).First(ctx)
	// fmt.Printf("user = %+v\n", user1)

	// user2, err := gorm.G[User](db).Take(ctx)
	// fmt.Printf("user1 = %+v\n", user2)

	// user3, err := gorm.G[User](db).Last(ctx)
	// fmt.Printf("user = %+v\n", user3)

	// errors.Is(err, gorm.ErrRecordNotFound)

	// user4 := db.First(&User{})
	// db.Take(&User{})
	// db.Last(&User{})

	// result := db.First(&user4)
	//result.RowsAffected // 返回找到的记录数
	//result.Error        // returns error or nil
	// 检查 ErrRecordNotFound 错误
	// errors.Is(result.Error, gorm.ErrRecordNotFound)

	var user5 = User{ID: 10}
	db.First(&user5)
	fmt.Printf("user = %+v\n", user5)

}
