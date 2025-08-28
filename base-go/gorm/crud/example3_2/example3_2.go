package example3_2

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	CreditCard []CreditCard
}

// type User struct {
// 	gorm.Model
// 	Name       string     `gorm:"unique;index;size:50"`
// 	CreditCard CreditCard `gorm:"foreignKey:UserName;references:Name"`
// }

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
}

func Run(db *gorm.DB) {
	fmt.Println("-------- example3_2 --------")
	db.AutoMigrate(&User{})
	db.AutoMigrate(&CreditCard{})

	// user := User{}
	// db.Create(&user)

	// card := CreditCard{Number: "Jinzhu", UserID: 1}
	// db.Create(&card)

	// card2 := CreditCard{Number: "Jinzhu2", UserID: 1}
	// db.Create(&card2)

	// user := User{}
	// db.Preload("CreditCards").First(&user, 1)
	// fmt.Println(user)

}
