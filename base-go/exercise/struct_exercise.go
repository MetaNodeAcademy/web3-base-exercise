package exercise

import (
	"fmt"
)

var noName4 = struct{}{}

var noNameStructChannel = make(chan struct{}, 0)

// 1.5 结构体
type Person struct {
	Name  string            `json:"name" gorm:"column:name"`
	Age   int               `json:"age" gorm:"column:age"`
	Call  func()            `json:"-" gorm:"column:call"`
	Map   map[string]string `json:"map" gorm:"column:map"`
	Ch    chan string       `json:"-" gorm:"column:ch"`
	Arr   [32]uint8         `json:"arr" gorm:"column:arr"`
	Slice []interface{}     `json:"slice" gorm:"column:slice"`
	Ptr   *int              `json:"-"`
	//0 Other `json:"-"`
}

type Custom struct {
	field1, field2, field3 int
	Slice                  []interface{}
}

type Custom2 struct {
	int
	string
	Other string
}

type Other struct {
}

type A struct {
	a string
}

func (a A) string() string {
	return a.a
}

func (a A) stringA() string {
	return a.a
}

func (a A) setA(v string) {
	a.a = v
}

func (a *A) stringPA() string {
	return a.a
}

func (a *A) setPA(v string) {
	a.a = v
}

type B struct {
	A
	b string
}

func (b B) string() string {
	return b.b
}

func (b B) stringB() string {
	return b.b
}

type C struct {
	B
	a string
	b string
	c string
	d []byte
}

func (c C) string() string {
	return c.c
}

func (c C) modityD() {
	c.d[2] = 33
}

func callStructMethod() {
	var a A
	a = A{
		a: "a",
	}
	a.string()
}

func NewC() C {
	return C{
		B: B{
			A: A{
				a: "a",
			},
			b: "b",
		},
		a: "ca",
		b: "cb",
		c: "cc",
		d: []byte{1, 2, 3},
	}
}

// 结构体
func StructExercise() {
	fmt.Println("--------结构体--------")
}

// 1.5.2 定义匿名结构体
func DefineAnonymousStruct() {
	fmt.Println("-----定义匿名结构体-----")
	var anonymousStruct = struct {
		Name string
		Age  int
		int
		string
	}{
		Name:   "Anonymous",
		Age:    30,
		int:    123,
		string: "Hello",
	}
	fmt.Println(anonymousStruct)
	fmt.Println(anonymousStruct.Name)
	fmt.Println(anonymousStruct.Age)
	fmt.Println(anonymousStruct.int)
	fmt.Println(anonymousStruct.string)
}

// 1.5.3 嵌套结构体
// func NestedStruct() {
// 	fmt.Println("-----嵌套结构体-----")
// 	var a A = A{a: "a"}
// 	var b B = B{A: a, b: "b"}
// 	var c C = C{A: a, B: b, a: "ca", b: "cb", c: "cc"}
// 	var d D = D{A: a, da: a, db: b}

// 	fmt.Println(c)
// 	fmt.Println(d.a)
// 	fmt.Println(d.da.a)
// }

// type A struct {
// 	a string
// }

// type B struct {
// 	A
// 	b string
// }

// type C struct {
// 	A
// 	B
// 	a string
// 	b string
// 	c string
// }

// type D struct {
// 	A
// 	da A
// 	db B
// }

// 1.5.4 定义结构体方法
func StructMethod() {
	var c C = NewC()
	var cp *C = &c
	fmt.Println(c.string())
	fmt.Println(c.stringA())
	fmt.Println(c.B.A.stringA())
	fmt.Println(c.stringB())
	fmt.Println(c.B.stringB())

	fmt.Println(cp.string())
	fmt.Println(cp.stringA())
	fmt.Println(cp.B.A.stringA())
	fmt.Println(cp.stringB())
	fmt.Println(cp.B.stringB())

	c.setA("new a")
	fmt.Println("After setA: ")
	fmt.Println(c.A.a)
	fmt.Println(cp.A.a)

	cp.setA("new a from pointer")
	fmt.Println("After setA from pointer: ")
	fmt.Println(c.A.a)
	fmt.Println(cp.A.a)

	c.setPA("new a from pointer to struct")
	fmt.Println("After setPA: ")
	fmt.Println(c.A.a)
	fmt.Println(cp.A.a)

	cp.setPA("new a from pointer to pointer")
	fmt.Println("After setPA from pointer to pointer: ")
	fmt.Println(c.A.a)
	fmt.Println(cp.A.a)

	cp.modityD()
	fmt.Println("After modifyD: ")
	fmt.Println(cp.d)

}
