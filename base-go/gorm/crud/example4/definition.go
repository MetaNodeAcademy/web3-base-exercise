package example4

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type strArr []string

type UserD struct {
	ID     int
	Skills strArr
}

// 存库
func (arr strArr) Value() (driver.Value, error) {
	return strings.Join(arr, ","), nil
}

// 读库
func (arr *strArr) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("Scan error on field type strArr")
	}
	*arr = strings.Split(string(bytes), ",")
	return nil
}

func RunDefinition(db *gorm.DB) {
	db.AutoMigrate(&UserD{})
	//db.Create(&UserD{Skills: strArr{"C++", "Python", "Go"}})

	var user UserD
	// db.First(&user, 1)
	// db.First(&user, "id = 1")
	db.Where("id = ?", 1).First(&user)
	fmt.Println(user)

	// sql.NullString{}
	// sql.NullTime{}

}
