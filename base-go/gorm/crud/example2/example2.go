package example2

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(255)"`
	Age      int
	Birthday time.Time
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("BeforeCreate")
	//u.Name = u.Name + "_123123"
	return
}

type CreditCard struct {
	gorm.Model
	Number     string
	BackUserId uint
}

type BankUser struct {
	gorm.Model
	Name string
	// CreditCard CreditCard `gorm:"embedded"`
}

func Run(db *gorm.DB) {
	fmt.Println("--------CRUD example2--------")
	// db.AutoMigrate(&User{})
	// db.AutoMigrate(&BankUser{})
	// db.AutoMigrate(&CreditCard{})

	// user := User{Name: "Jinzhu2", Age: 19, Birthday: time.Now()}
	// 通过数据库指针创建
	// result := db.Create(&user)
	// fmt.Printf("result = %+v 影响行数: %d\n", result, result.RowsAffected)

	// users := []User{
	// 	{Name: "Jinzhu1", Age: 18, Birthday: time.Now()},
	// 	{Name: "Jinzhu2", Age: 19, Birthday: time.Now()},
	// 	{Name: "Jinzhu3", Age: 20, Birthday: time.Now()},
	// }
	// db.Create(&users)
	// fmt.Printf("users = %+v\n", users)

	// var user User
	// db.Debug().First(&user)
	// fmt.Printf("user = %+v\n", user)

	// user.ID = 10
	// db.Debug().First(&user)
	// db.Debug().First(&user, []int{7, 8, 9})
	// fmt.Printf("user = %+v\n", user)

	var user User
	user.ID = 1
	db.Debug().First(&user, "name = ?", "测试")
	fmt.Printf("user = %+v\n", user)
}
