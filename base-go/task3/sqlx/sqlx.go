package sqlx

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Employee struct {
	Id         int    `db:"id"`
	Name       string `db:"name"`
	Department string `db:"department"`
	Salary     int    `db:"salary"`
}

type Book struct {
	ID     int    `db:"id"`
	Title  string `db:"title"`
	Author string `db:"author"`
	Price  int    `db:"price"`
}

func SqlxExercise() {
	fmt.Println("--------sqlx--------")
	db := InitDB()
	// 题目1
	// exampleSqlx1(db)
	exampleSqlx2(db)
}

// 题目1
func exampleSqlx1(db *sqlx.DB) {
	/*
		题目1：使用SQL扩展库进行查询
		假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
		要求 ：
		编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
		编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
	*/
	var employees []Employee
	sql := "SELECT * FROM employees WHERE department = :department ORDER BY id DESC"
	param := map[string]interface{}{
		"department": "技术部",
	}

	smt, err := db.PrepareNamed(sql)
	if err != nil {
		log.Fatalln("PrepareNamed failed:", err)
	}
	defer smt.Close()
	selectErr := smt.Select(&employees, param)
	if selectErr != nil {
		log.Fatalln(err)
	}
	log.Printf("SQL: %s \nemployees: %+v\n", sql, employees)

	// 编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
	sql2 := "SELECT * FROM employees ORDER BY salary DESC LIMIT 1"
	var employee Employee
	err2 := db.Get(&employee, sql2)
	if err2 != nil {
		log.Fatalln("Get failed:", err2)
	}
	log.Printf("SQL: %s \nemployee: %+v\n", sql2, employee)
}

// 题目2
func exampleSqlx2(db *sqlx.DB) {
	/*
		题目2：实现类型安全映射
		假设有一个 books 表，包含字段 id、title、author、price
		要求 ：
		定义一个 Book 结构体，包含与 books 表对应的字段。
		编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
	*/
	var books []Book
	db.Select(&books, "SELECT * FROM books WHERE price > ?", 100)
	log.Printf("books: %v", books)
}

func InitDB() *sqlx.DB {
	// DSN 格式：用户名:密码@tcp(主机:端口)/数据库名?参数
	dsn := "root:Asdf123#@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True"
	// 连接数据库
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalln("连接数据库失败:", err)
	}
	// defer db.Close()
	log.Printf("数据库连接成功 db: %+v\n", db)
	return db
}
