package exercise

import "fmt"

// 全局变量
var gs string = "Hello, world!"
var gzero int
var gb2 bool = true
var gm map[string]int
var garray [2]byte
var p *int

var (
	vi int = 123
	vb bool
	vs = "test"
)

var (
	group1      = 2
	group2 byte = 2
)

// 多变量定义
var ma, mb, bc int = 1, 2, 3
var me, mf, mg int
var mh, mi, mj = 1, 2, "mj"

// 1.2 基本数据类型
func BaseDataType() {
	// 定义基本数据类型
	fmt.Println("-----定义基本数据类型-----")
	// 十六进制
	var a uint8 = 0xF
	var b uint8 = 0xf
	// 八进制
	var c uint8 = 017
	var d uint8 = 0o17
	var e uint8 = 0o17
	// 二进制
	var f uint8 = 0b1111
	var g uint8 = 0b1111
	// 十进制
	var h uint8 = 15
	// 检查是否一致
	fmt.Println(a == b)
	fmt.Println(b == c)
	fmt.Println(c == d)
	fmt.Println(d == e)
	fmt.Println(e == f)
	fmt.Println(f == g)
	fmt.Println(g == h)
	fmt.Println(h == a)
	// 浮点
	var float1 float32 = 10
	float2 := 10.0
	float1 = float32(float2)
	fmt.Println(float1)
	fmt.Println(float1 == float32(float2))
	// 复数
	var c1 complex64
	c1 = 1.10 + 0.1i
	c2 := 1.10 + 0.1i
	c3 := complex(1.10, 0.1)
	fmt.Println(c1 == complex64(c2))
	fmt.Println(complex128(c1) == c2)
	fmt.Println(c2 == c3)
	x := real(c2)
	y := imag(c2)
	fmt.Println(x)
	fmt.Println(y)
	// bytes
	var s string = "Hello, world!"
	var bytes []byte = []byte(s)
	fmt.Println("convert \"Hello, world!\" to bytes:", bytes)
	fmt.Println(string(bytes) == s)
	// bytes 转回 string
	var bytes2 []byte = []byte{72, 101, 108, 108, 111, 44, 32, 119, 111, 114, 108, 100, 33}
	var s2 string = string(bytes2)
	fmt.Println("convert bytes back to string:", s2)
	fmt.Println(s2 == s)
	// rune
	var r1 rune = 'a'
	var r2 rune = '世'
	fmt.Println(r1)
	fmt.Println(r2)
	var s3 string = "abc,你好,世界!"
	var runes []rune = []rune(s3)
	fmt.Println(runes)
	fmt.Println(len(runes))
	// 字符串
	var s4 string = "Hello\nworld!\n"
	fmt.Println(s4)
	var s5 string = `Hello
world!
`
	fmt.Println(s5)
	fmt.Println(s4 == s5)
	fmt.Println("bytes runes 关系")
	var s6 string = "Go语言👍"
	var bytes3 []byte = []byte(s6)
	var runes3 []rune = []rune(s6)
	fmt.Println("string length", len(s6))
	fmt.Println("bytes length", len(bytes3))
	fmt.Println("runes length", len(runes3))
	fmt.Println(s6[0:12])
	fmt.Println(string(bytes3[0:12]))
	fmt.Println(string(runes3[0:5]))
	// 布尔
	var b1 bool = true
	var b2 bool = false
	fmt.Println("布尔值 b1:", b1)
	fmt.Println("布尔值 b2:", b2)
	fmt.Println("布尔值 b1 == b2:", b1 == b2)
	// 零值
	var digit int
	var s7 string
	var b3 bool
	var f4 float64
	var c4 complex128
	fmt.Println("int zero value:", digit)
	fmt.Println("string zero value:", s7)
	fmt.Println("bool zero value:", b3)
	fmt.Println("float64 zero value:", f4)
	fmt.Println("complex128 zero value:", c4)
}

// 1.3 定义变量
func DefineVariable() {
	// 全局变量
	fmt.Println("-----全局变量-----")
	fmt.Println("global string:", gs)
	fmt.Println("global int zero value:", gzero)
	fmt.Println("global bool:", gb)
	fmt.Println("global map:", gm)
	fmt.Println("global array:", garray)
	fmt.Println("global pointer:", p)
	fmt.Println("vi:", vi)
	fmt.Println("vb:", vb)
	fmt.Println("vs:", vs)
	fmt.Println("group1:", group1)
	fmt.Println("group2:", group2)

	gm = make(map[string]int, 0)
	gm["one"] = 1
	gm["two"] = 2
	gm["three"] = 3
	fmt.Println("global map:", gm)

	// 局部变量
	fmt.Println("-----局部变量-----")
	method1()
	ma1, mb1 := method2()
	ma2, mb2 := method3()
	ma3, mb3 := method4()
	fmt.Println("method2 return values:", ma1, mb1)
	fmt.Println("method3 return values:", ma2, mb2)
	fmt.Println("method4 return values:", ma3, mb3)
	// 多变量定义
	fmt.Println(ma, mb, bc, me, mf, mg, mh, mi, mj)
	methodMultivariable()
}

func method1() {
	// 方式1: 类型推导
	a := 1
	// 方式2: 显式类型声明
	var b int = 2
	// 方式3: 仅声明不赋值
	var c int
	fmt.Println(a, b, c)
}

func method2() (a int, b string) {
	// 必须声明return关键字
	return 1, "method2 hello"
}

func method3() (a int, b string) {
	a = 1
	b = "hello"
	return
}

func method4() (a int, b string) {
	return
}

// 多变量定义
func methodMultivariable() {
	var k, l, m int = 1, 2, 3
	var n, o, p int
	q, r, s := 4, 5, "test"
	t, u, v := 6, 7, "vvv"
	fmt.Println(k, l, m, n, o, p, q, r, s, t, u, v)
}