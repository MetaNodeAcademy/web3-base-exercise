package exercise

import (
	"fmt"
	"time"
)

var ga int
var gb int = 10

// 1.11 变量的作用域
func DefineScope() {
	fmt.Println("-----变量的作用域-----")
	var a int
	if b := 1; b == 0 {
		fmt.Println("b == 0")
	} else {
		c := 2
		fmt.Println("declare c = ", c)
		fmt.Println("b == 1 ")
	}
	//fmt.Println(b)
	//fmt.Println(c)

	switch d := 3; d {
	case 1:
		e := 4
		fmt.Println("declare e = ", e)
		fmt.Println("d == 1")
	case 3:
		f := 4
		fmt.Println("declare f = ", f)
		fmt.Println("d == 3")
	}
	//fmt.Println(e)
	//fmt.Println(f)

	for i := 0; i < 1; i++ {
		forA := 1
		fmt.Println("declare forA = ", forA)
	}

	select {
	case <-time.After(time.Second * 3):
		selectA := 1
		fmt.Println("declare selectA = ", selectA)
	}

	// 匿名代码块
	{
		blockA := 1
		fmt.Println("declare blockA = ", blockA)
	}
	//fmt.Println("declare blockA = ", blockA)
	fmt.Println("declare a = ", a)

	// 全局变量
	{
		fmt.Println("global variable, ga = ", ga)
		ga = 3
		fmt.Println("global variable, ga = ", ga)
		ga := 10
		fmt.Println("local variable, ga = ", ga)
		ga--
		fmt.Println("local variable, ga = ", ga)
	}

	// 作用域优先级
	var gb int = 4
	fmt.Println("local variable, gb = ", gb)
	if gb := 3; gb == 3 {
		fmt.Println("if gb = ", gb)
		gb--
		fmt.Println("if gb = ", gb)
	}
	fmt.Println("local variable, gb = ", gb)
	fmt.Println("global variable, gb = ", testGlobalVariable())

}

func testGlobalVariable() int {
	return gb
}
