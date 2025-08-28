package example4

import (
	"fmt"

	"gorm.io/gorm"
)

type Parent struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

func (p *Parent) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("BeforeCreate Parent")
	return
}

func (p *Parent) AfterCreate(tx *gorm.DB) (err error) {
	fmt.Println("AfterCreate Parent")
	return
}

type Child struct {
	Parent
	Age int
}

func (c *Child) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("BeforeCreate Child")
	return
}

func (c *Child) AfterCreate(tx *gorm.DB) (err error) {
	fmt.Println("AfterCreate Child")
	return
}

func RunHook(db *gorm.DB) {
	db.AutoMigrate(&Parent{}, &Child{})
	db.Create(&Child{Parent: Parent{Name: "fu"}, Age: 10})
}
