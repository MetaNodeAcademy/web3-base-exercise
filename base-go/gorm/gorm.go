package gorm

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type Parent struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type Child struct {
	Parent
	Age int
}

// 初始化数据库链接
func InitDB(dst ...interface{}) *gorm.DB {
	// 日志输出
	logger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,       // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)
	db, err := gorm.Open(mysql.Open("root:Asdf123#@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
		Logger:                                   logger, // 日志
		DryRun:                                   false,  // 是否测试SQL模式
		PrepareStmt:                              true,   // 是否预编译
		SkipDefaultTransaction:                   false,  // 是否跳过默认事务
		DisableNestedTransaction:                 false,  // 是否禁用嵌套事务
		AllowGlobalUpdate:                        false,  // 是否允许全局更新
		DisableAutomaticPing:                     false,  // 是否关闭 GORM 启动时自动 ping 数据库
		DisableForeignKeyConstraintWhenMigrating: false,  // 是否创建数据库层面的外键约束
		CreateBatchSize:                          1000,   // 批量创建数据时，一次创建多少条数据
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false, // 是否使用单数表名
		},
	})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	// 最大打开连接数
	sqlDB.SetMaxOpenConns(10)
	// 最大空闲连接数
	sqlDB.SetMaxIdleConns(10)
	// 连接最大生命周期
	sqlDB.SetConnMaxLifetime(time.Hour * 1)
	// 设置全局 string -> varchar(255)
	db.Config.Dialector.(*mysql.Dialector).DefaultStringSize = 255
	// 设置打印SQL结果
	db.Callback().Query().After("gorm:after_query").Register("print_result", func(db *gorm.DB) {
		if db.Statement.Dest != nil {
			fmt.Printf("result: %+v\n", db.Statement.Dest)
		}
	})
	amErr := db.AutoMigrate(dst...)
	if amErr != nil {
		panic(amErr)
	}
	fmt.Printf("数据库连接成功 db: %+v\n", db.Config.Name())
	return db
}

func GormExercise() {
	fmt.Println("--------Gorm--------")
	// 创建链接
	db := InitDB()
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()
	// 示例
	//example1.Run(db)
	//example2.Run(db)
	//example3.Run(db)
	//example3_2.Run(db)
	//example3_3.Run(db)
	//example4.Run(db)

	//db.Debug().Where("id = ?", 2).Delete(&Child{})

	fmt.Println("--------Gorm结束--------")
}
