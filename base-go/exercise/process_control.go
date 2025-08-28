package exercise

import "fmt"

type CustomType struct {
}

// 流程控制
func ProcessControl() {
	fmt.Println("-----流程控制-----")
	var a int = 10
	if b := 1; a > 10 && a < 20 {
		b = 2
		// c = 2
		fmt.Println("a > 10")
	} else if c := 3; b > 1 {
		b = 3
		fmt.Println("b > 1")
	} else {
		fmt.Println("else")
		if c == 3 {
			fmt.Println("c == 3")
		}
		fmt.Println("b = ", b)
		fmt.Println("c = ", c)
	}
}

// swtich case
func SwitchCase() {
	fmt.Println("-----swtich case-----")
	// 基本用法
	a := "Hello World"
	switch a {
	case "Hello World":
		fmt.Println("Hello World")
	case "Hello", "World":
		fmt.Println("Hello")
	default:
		fmt.Println("default")
	}
	// 变量b仅在当前Switch代码块内有效
	switch b := 5; b {
	case 1:
		fmt.Println("b == 1")
	case 2:
		fmt.Println("b == 2")
	case 3, 4:
		fmt.Println("b == 3 or b == 4")
	case 5:
		fmt.Println("b == 5")
	default:
		fmt.Println("default b = ", b)
	}
	// 不指定判断变量，直接在case中添加判定条件
	b := 5
	switch {
	case a == "H":
		fmt.Println("a == H")
	case b == 3:
		fmt.Println("b == 3")
	case b == 5, a == "Hello World":
		fmt.Println("a == Hello World or b == 5")
	default:
		fmt.Println("default case")
	}
	// 判断类型
	var d interface{}
	//e := byte(1)
	//e := "Hello World"
	//d = &e
	d = "Hello World"
	switch t := d.(type) {
	case byte:
		fmt.Println("d is byte", t)
	case *byte:
		fmt.Println("d is byte point", t)
	case *int:
		fmt.Println("d is int point", t)
	case string:
		fmt.Println("d is string", t)
	case *string:
		fmt.Println("d is string point", t)
	case CustomType:
		fmt.Println("d is CustomType", t)
	case *CustomType:
		fmt.Println("d is CustomType point", t)
	default:
		fmt.Println("d is unknown type", t)
	}
}
