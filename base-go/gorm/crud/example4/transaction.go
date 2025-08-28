package example4

import (
	"gorm.io/gorm"
)

type User struct {
	ID   int
	Name string
	Age  int
}

func RunTransaction(db *gorm.DB) {
	db.AutoMigrate(&User{})

	// var user User
	// db.Create(&User{Name: "u1", Age: 10})
	// db.First(&user)
	// fmt.Println(user)

	// 禁用事物
	// tx := db.Session(&gorm.Session{SkipDefaultTransaction: true})
	// tx.First(&user, 1)
	// tx.Model(&user).Update("Age", 19)

	// db.Transaction(func(tx *gorm.DB) error {
	// 	// 在事物中执行 db 操作, 要使用 tx
	// 	if err := tx.Create(&User{Name: "u2", Age: 20}).Error; err != nil {
	// 		// 返回任务错误会回滚
	// 		return err
	// 	}

	// 	if err := tx.Create(&User{Name: "u3", Age: 21}).Error; err != nil {
	// 		return err
	// 	}
	// 	// 返回 nil 提交事务
	// 	return nil
	// 	//return errors.New("rollback")
	// })

	// 嵌套事务
	// db.Transaction(func(tx *gorm.DB) error {
	// 	tx.Create(&User{Name: "u1", Age: 18})
	// 	// 将会回滚
	// 	tx.Transaction(func(itx *gorm.DB) error {
	// 		itx.Create(&User{Name: "u2", Age: 19})
	// 		return errors.New("rollback user u2")
	// 	})
	// 	// 上面回滚不影响
	// 	tx.Transaction(func(itx *gorm.DB) error {
	// 		itx.Create(&User{Name: "u3", Age: 20})
	// 		return nil
	// 	})
	// 	return nil
	// })

	// 手动事务
	// 开始事务
	// tx := db.Begin()
	// 执行 db 操作, 需使用 tx
	// tx.Create(&User{Name: "u1", Age: 18})
	// tx.Create(&User{Name: "u2", Age: 19})
	// 遇到错误时回滚事务
	//tx.Rollback()
	// 提交事务
	// tx.Commit()

	// tx := db.Begin()
	// tx.Create(&User{Name: "this is sp1", Age: 18})
	// tx.SavePoint("sp1")
	// tx.Create(&User{Name: "zhangsan"})
	// tx.RollbackTo("sp1")
	// tx.Commit()

	tx := db.Begin()
	tx.Set("gorm:table_options", "ENGINE=InnoDB")
}
