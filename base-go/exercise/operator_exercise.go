package exercise

import (
	"fmt"
)

// 运算操作符
func ArithmeticOperator() {
	fmt.Println("-----运算操作符-----")
	a, b := 1, 2
	fmt.Println("a + b =", a+b) // 加法
	fmt.Println("a - b =", a-b) // 减法
	fmt.Println("a * b =", a*b) // 乘法
	fmt.Println("a / b =", a/b) // 除法
	fmt.Println("a % b =", a%b) // 取模
	// 自增自减
	a++
	a--
	fmt.Println("a++ =", a)
	fmt.Println("a-- =", a)
	// 错误格式
	//++a
	//--a
	//b := a++ +1
	//c := a--
	// 混合计算
	a1 := 10 + 0.1
	b1 := byte(1) + 1
	fmt.Println(a1)
	fmt.Println(b1)
	sum := a1 + float64(b1)
	fmt.Println(sum)
	sub := byte(a1) - b1
	fmt.Println(sub)
	mul := a1 * float64(b1)
	fmt.Println(mul)
	div := byte(a1) / b1
	fmt.Println(div)
	// 关系运算符
	a2 := 1
	b2 := 5
	fmt.Println(a2 == b2) // 等于
	fmt.Println(a2 != b2) // 不等于
	fmt.Println(a2 > b2)  // 大于
	fmt.Println(a2 < b2)  // 小于
	fmt.Println(a2 >= b2) // 大于等于
	fmt.Println(a2 <= b2) // 小于等于
	// 逻辑运算符
	a3 := true
	b3 := false
	fmt.Println(a3 && b3)    // 逻辑与
	fmt.Println(a3 || b3)    // 逻辑或
	fmt.Println(!(a3 && b3)) // 逻辑非
	// 位运算符
	fmt.Println(0 & 0) // 按位与
	fmt.Println(0 | 0) // 按位或
	fmt.Println(0 ^ 0) // 按位异或
	fmt.Println(0 & 1) // 按位与
	fmt.Println(0 | 1) // 按位或
	fmt.Println(0 ^ 1) // 按位异或
	fmt.Println(1 & 1) // 按位与
	fmt.Println(1 | 1) // 按位或
	fmt.Println(1 ^ 1) // 按位异或
	fmt.Println(1 & 0) // 按位与
	fmt.Println(1 | 0) // 按位或
	fmt.Println(1 ^ 0) // 按位异或
	// 赋值运算符
	fmt.Println("------赋值运算符------")
	a4, b4 := 1, 2
	var c4 int
	c4 = a4 + b4
	fmt.Println(c4)
	plusAssignment(c4, a4)
	subAssignment(c4, a4)
	mulAssignment(c4, a4)
	divAssignment(c4, a4)
	modAssignment(c4, a4)
	leftMoveAssignment(c4, a4)
	rightMoveAssignment(c4, a4)
	andAssignment(c4, a4)
	orAssignment(c4, a4)
	norAssignment(c4, a4)
	// 其他运算符
	fmt.Println("------其他运算符------")
	a5 := 4
	var ptr *int
	fmt.Println(a5)
	ptr = &a5
	fmt.Println(*ptr)
	// 运算符优先级
	priority()
}

// 运算符优先级
func priority() {
	fmt.Println("------运算符优先级------")
	var a int = 21
	var b int = 10
	var c int = 16
	var d int = 5
	var e int

	e = (a + b) * c / d // 31 * 16 / 5
	fmt.Println("(a + b) * c / d = ", e)
	e = ((a + b) * c) / d // (31 * 16) / 5
	fmt.Println("((a + b) * c) / d = ", e)
	e = (a + b) * (c / d) // 31 * (16 / 5)
	fmt.Println("(a + b) * (c / d) = ", e)
	e = a + (b*c)/d // 21 + (10 * 16) / 5
	fmt.Println("a + (b * c) / d = ", e)

	f := 3 + 4 ^ 3 | 2&2*3<<1
	fmt.Println(f == 12)

}

func plusAssignment(c, a int) {
	c += a
	fmt.Println(c)
}

func subAssignment(c, a int) {
	c -= a
	fmt.Println(c)
}

func mulAssignment(c, a int) {
	c *= a
	fmt.Println(c)
}

func divAssignment(c, a int) {
	c /= a
	fmt.Println(c)
}

func modAssignment(c, a int) {
	c %= a
	fmt.Println(c)
}

func leftMoveAssignment(c, a int) {
	c <<= a
	fmt.Println(c)
}

func rightMoveAssignment(c, a int) {
	c >>= a
	fmt.Println(c)
}

func andAssignment(c, a int) {
	c &= a
	fmt.Println(c)
}

func orAssignment(c, a int) {
	c |= a
	fmt.Println(c)
}

func norAssignment(c, a int) {
	c ^= a
	fmt.Println(c)
}
