package exercise

import (
	"fmt"
	"strconv"
)

type Supplier interface {
	Get() string
}

type DigitSupplier struct {
	value int
}

func (i *DigitSupplier) Get() string {
	return fmt.Sprintf("%d", i.value)
}

// 1.16 类型转换
func TypeConversion() {
	fmt.Println("-------- 1.16 类型转换--------")
	var i int32 = 17
	var b byte = 5
	var f float32
	// 数字类型可以直接强转
	f = float32(i) / float32(b)
	fmt.Printf("f 的值为: %f\n", f)
	// 当int32类型强转成byte时, 高位被直接舍弃
	var i2 int32 = 256
	var b2 byte = byte(i2)
	fmt.Printf("b2 的值为: %d\n", b2)

	str := "Hello, 123, 你好"
	var bytes []byte = []byte(str)
	var runes []rune = []rune(str)
	fmt.Printf("bytes 的值为: %v\n", bytes)
	fmt.Printf("runes 的值为: %v\n", runes)

	str2 := string(bytes)
	str3 := string(runes)
	fmt.Printf("str2 的值为: %s\n", str2)
	fmt.Printf("str3 的值为: %s\n", str3)

}

// 字符串和数字转换
func StringToNumber() {
	fmt.Println("--------字符串和数字转换--------")
	str := "123"
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	fmt.Printf("字符串转换为 int 的值为: %d\n", num)
	str1 := strconv.Itoa(num)
	fmt.Printf("int 转换为字符串的值为: %s\n", str1)

	ui64, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		panic(err)
	}
	fmt.Printf("字符串转换为 uint64 的值为: %d\n", ui64)

	str2 := strconv.FormatUint(ui64, 2)
	fmt.Printf("uint64 转换为字符串的值为: %s\n", str2)
}

// 1.16.3 接口类型转换
func InterfaceTypeConversion() {
	fmt.Println("--------接口类型转换--------")
	var i interface{} = 3
	a, ok := i.(int)
	if ok {
		fmt.Printf("i 的值为: %d\n", a)
	} else {
		fmt.Println("i 的类型不是 int")
	}
	switch v := i.(type) {
	case int:
		fmt.Printf("i is a int %d\n", v)
	case string:
		fmt.Printf("i is a string %s\n", v)
	default:
		fmt.Println("i is unknown type %s\n", v)
	}

	var aa Supplier = &DigitSupplier{value: 1}
	fmt.Println(aa.Get())
	b, ok := (aa).(*DigitSupplier)
	fmt.Println(b, ok)
}

type SameFieldA struct {
	name  string
	value int
}

type SameFieldB struct {
	name  string
	value int
}

func (s *SameFieldB) getValue() int {
	return s.value
}

// 1.16.4 结构体类型转换
func StructTypeConversion() {
	fmt.Println("-----结构体类型转换-----")
	a := SameFieldA{
		name:  "a",
		value: 1,
	}
	b := SameFieldB(a)
	fmt.Printf("convert a to b: %d\n", b.getValue())
	// 只能结构体类型实例之间相互转换，指针不可以相互转换
	var c interface{} = &a
	_, ok := c.(*SameFieldB)
	fmt.Printf("c is *SameFieldB: %v \n", ok)
}
