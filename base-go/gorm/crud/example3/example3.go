package example3

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string
	CompanyID int
}

type UserB struct {
	gorm.Model
	Name       string
	CompanyIDS int
	Company    Company `gorm:"foreignkey:CompanyIDS"` // users.CompanyIDS = company.ID
}

type UserC struct {
	gorm.Model
	Name      string
	CompanyID int
	//Company   Company `gorm:"references:CardNumber"` // users.CompanyID = company.CardNumber
	Company Company `gorm:"foreignkey:CompanyID;references:CardNumber"` // users.CompanyID = company.CardNumber
}

type Company struct {
	ID         int
	CardNumber int `gorm:"unique"`
	Name       string
}

func Run(db *gorm.DB) {
	fmt.Println("-------- example3 --------")
	db.AutoMigrate(&User{})
	// db.AutoMigrate(&UserB{})
	// db.AutoMigrate(&UserC{})
	db.AutoMigrate(&Company{})

	// user := User{Name: "user normal", CompanyID: 1}
	// db.Create(&user)

	// company := Company{Name: "John com", CardNumber: 12345678}
	// db.Create(&company)

	// user := User{}
	// db.Last(&user)
	// fmt.Println(user)

	// user := UserB{CompanyIDS: 1, Name: "user_reference_id"}
	// db.Create(&user)

	// user := User{}
	// db.Preload("Company").First(&user)
	// fmt.Println(user)

	// user := UserC{CompanyID: 12333, Name: "user_ccc"}
	// db.Create(&user)

}
