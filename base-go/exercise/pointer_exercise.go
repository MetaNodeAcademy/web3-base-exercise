package exercise

import (
	"fmt"
	"unsafe"
)

// 指针
var i2 int
var p1 *int = &i2
var p2 *string = &gs

// 1.4 指针
func Pointer() {
	// TODO
	fmt.Println("-----指针-----")
	var i int = 1
	var s string = "Hello"
	p1 = &i
	p2 = &s
	var p3 **string = &p2
	var p4 *byte
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)
	fmt.Println(p4)
	fmt.Println(*p1)
	fmt.Println(*p2)
	fmt.Println(**p3)
	// p4的内存地址
	fmt.Println("p4 address:", &p4)
}

// 1.4.3 修改指针指向的值
func ModifyPointer() {
	fmt.Println("-----修改指针指向的值-----")
	var a int = 2
	var p *int
	p = &a
	fmt.Println(p, &a)
	var pp **int
	pp = &p
	fmt.Println(pp, p)
	// 改修指针指向的值
	**pp = 3
	fmt.Println(pp, *pp, p)
	fmt.Println(**pp, *p, a)
	fmt.Println(a, &a)
}

// 1.4.4 unsafe unitptr
func UnsafeUnitPtr() {
	fmt.Println("-----unsafe unitptr-----")
	var a string = "Hello, world!"
	upA := uintptr(unsafe.Pointer(&a))
	upA += 6

	c := (*uint8)(unsafe.Pointer(upA))
	fmt.Println(*c)
	*c = 12
	fmt.Println(*c) // 输出 "Hello, Gorld!"
}
