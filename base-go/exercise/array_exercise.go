package exercise

import (
	"fmt"
)

type Custome struct {
	a int
}

// 1.12 数组
func ArrayExercise() {
	fmt.Println("-------数组-------")
	// 仅声明
	var a [5]int
	fmt.Println("a = ", a)

	var marr [2]map[string]string
	fmt.Println("marr = ", marr)
	// map的零值是nil，虽然打印出来是非空值，但真实的值是ni
	//marr[0]["test"] = "1"

	// 声明以及初始化
	var b [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("b = ", b)

	// 类型推导声明方式
	var c = [5]string{"c1", "c2", "c3", "c4", "c5"}
	fmt.Println("c = ", c)

	d := [3]int{3, 2, 1}
	fmt.Println("d = ", d)

	// 使用 ... 代替数组长度
	autoLen := [...]string{"auto1", "auto3", "auto4", "auto5"}
	fmt.Println("autoLen = ", autoLen)
	// 初始化时，元素个数不能超过数组声明的长度
	//overLen := [...][2]int{1, 2, 3}
}

// 1.12.2 访问数组
func ArrayAccess() {
	fmt.Println("-------访问数组-------")
	a := [5]int{5, 4, 3, 2, 1}
	// 方式1, 使用下标读取数据
	element := a[2]
	fmt.Println("element = ", element)
	// 方式2, 使用range遍历
	for i, v := range a {
		fmt.Printf("下标 %d, 值 %d\n", i, v)
	}
	for i := range a {
		fmt.Printf("下标 %d, 值 %d\n", i, a[i])
	}
	// 读取数组长度
	fmt.Println("len(a) = ", len(a))
	// 使用下标，for循环遍历数组
	for i := 0; i < len(a); i++ {
		fmt.Printf("下标 %d, 值 %d\n", i, a[i])
	}
}

// 1.12.3 多维数组
func MultiDimensionalArray() {
	fmt.Println("-------多维数组-------")
	// 二维数组
	a := [3][2]int{
		{0, 1},
		{2, 3},
		{4, 5},
	}
	fmt.Println("a = ", a)
	// 三维数组
	b := [3][2][2]int{
		{
			{0, 1},
			{2, 3},
		},
		{
			{4, 5},
			{6, 7},
		},
		{
			{8, 9},
			{10, 11},
		},
	}
	fmt.Println("b = ", b)
	// 也可以省略各个位置的初始化，在后续代码中赋值
	c := [3][3][3]int{}
	c[2][2][1] = 5
	c[1][2][1] = 4
	fmt.Println("c = ", c)
	// 多维数组访问
	d := [3][2][2]int{
		{{0, 1}, {2, 3}},
		{{4, 5}, {6, 7}},
		{{8, 9}, {10, 11}},
	}
	layer1 := d[0]
	layer2 := d[0][1]
	layer3 := d[0][1][1]
	fmt.Println("layer1 = ", layer1)
	fmt.Println("layer2 = ", layer2)
	fmt.Println("layer3 = ", layer3)
	// 多维数组遍历时，需要使用嵌套for循环遍历
	for i, v := range d {
		fmt.Printf("第%d层 %v\n", i, v)
		for j, inner := range v {
			fmt.Printf("第%d层第%d列 %v\n", i, j, inner)
			for k, innerInner := range inner {
				fmt.Printf("第%d层第%d列第%d行 %v\n", i, j, k, innerInner)
			}
		}
	}
}

// 1.12.4 数组作为参数
func ArrayAsParam() {
	// fmt.Println("-------多维数组-------")
	// a := [5]int{5, 4, 3, 2, 1}
	// fmt.Println("before a = ", a)

	// receiveArray(a)
	// fmt.Println("after a = ", a)

	// receiveArrayPointer(&a)
	// fmt.Println("after a = ", a)

	// i1 := 5
	// i2 := 4
	// i3 := 3
	// i4 := 2
	// i5 := 1
	// b := [5]*int{&i1, &i2, &i3, &i4, &i5}
	// fmt.Printf("before b = %v *param[1] = %d\n", b, *b[1])
	// receivePointerArray(b)
	// fmt.Printf("before b = %v *param[1] = %d\n", b, *b[1])

	c1 := Custome{a: 5}
	c2 := Custome{a: 4}
	c3 := Custome{a: 3}
	c4 := Custome{a: 2}
	c5 := Custome{a: 1}
	c := [5]Custome{c1, c2, c3, c4, c5}
	d := [5]*Custome{&c1, &c2, &c3, &c4, &c5}
	fmt.Printf("before c = %v c[1].a = %d\n", c, c[1].a)
	receiveCustomArray(c)
	fmt.Printf("before c = %v c[1].a = %d\n", c, c[1].a)

	fmt.Printf("before d = %v d[1].a = %d\n", c, d[1].a)
	receivePointerCustomArray(d)
	fmt.Printf("before d = %v d[1].a = %d\n", c, d[1].a)
}

func receiveArray(param [5]int) {
	fmt.Println("in receiveArray func, before param = ", param)
	param[1] = -5
	fmt.Println("in receiveArray func, after param = ", param)
}

func receiveArrayPointer(param *[5]int) {
	fmt.Println("in receiveArrayPointer func, before param = ", param)
	param[1] = -5
	fmt.Println("in receiveArrayPointer func, after param = ", param)
}

func receivePointerArray(param [5]*int) {
	fmt.Println("in receivePointerArray func, before param = ", param)
	*param[1] = -5
	fmt.Println("in receivePointerArray func, after param = ", param)
}

func receiveCustomArray(param [5]Custome) {
	fmt.Println("in receiveCustom func, before param = ", param)
	param[1].a = -6
	fmt.Println("in receiveCustom func, after param = ", param)
}

func receivePointerCustomArray(param [5]*Custome) {
	fmt.Println("in receivePointerCustomArray func, before param = ", param)
	(*param[1]).a = -7
	fmt.Println("in receivePointerCustomArray func, after param = ", param)
}