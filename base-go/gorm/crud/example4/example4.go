package example4

import (
	"fmt"

	"gorm.io/gorm"
)

type Dog struct {
	ID   int
	Name string
	Toy  Toy `gorm:"polymorphicType:Kind;polymorphicId:TID;polymorphicValue:dog"`
}

type Cat struct {
	ID   int
	Name string
	Age  int
	Toy  Toy `gorm:"polymorphicType:Kind;polymorphicId:TID;polymorphicValue:cat"`
}

func (c *Cat) AfterCreate(tx *gorm.DB) (err error) {
	fmt.Println("AfterCreate Cat")
	return nil
}

func (c *Cat) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("BeforeCreate Cat")
	return nil
}

type Toy struct {
	ID   int
	Name string
	Kind string
	TID  int
}

func Run(db *gorm.DB) {
	fmt.Println("--------example4--------")

	// db.AutoMigrate(&Cat{}, &Dog{}, &Toy{})

	// 多态
	// db.Create(&Dog{Name: "wangcai1", Toy: Toy{Name: "ball"}})
	// db.Create(&Dog{Name: "wangcai2", Toy: Toy{Name: "ball"}})

	// db.Create(&Cat{Name: "mimi1", Toy: Toy{Name: "mouse"}, Age: 1})
	// db.Create(&Cat{Name: "mimi2", Toy: Toy{Name: "mouse"}, Age: 2})
	// db.Create(&Cat{Name: "mimi3", Toy: Toy{Name: "mouse"}, Age: 3})

	// var dog Dog
	// var cat Cat
	// db.Preload("Toy").First(&dog)
	// db.Preload("Toy").First(&cat)
	// fmt.Println(dog, cat)

	//RunHook(db)
	//RunTransaction(db)
	//RunDefinition(db)

}
