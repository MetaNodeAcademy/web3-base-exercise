package exercise

// type A struct {
// 	i int
// }

// func (a *A) add(v int) int {
// 	a.i += v
// 	return a.i
// }

// // 声明函数变量
// var function1 func(int) int

// // 声明闭包
// var squart2 func(int) int = func(p int) int {
// 	p *= p
// 	return p
// }

// var customVar = customHasReturn()

// 函数
// func FunctionExercixe() {
// 	fmt.Println("-------FunctionExercixe-------")
// 	// 函数
// 	custom()
// 	// 闭包
// 	a := A{1}
// 	function1 = a.add
// 	// 声明一个闭包并直接执行
// 	// 此闭包返回值是另外一个闭包
// 	// 返回值闭包赋值给returnFunc变量
// 	returnFunc := func() func(int, string) (int, string) {
// 		fmt.Println("this is a closure")
// 		return func(i int, s string) (int, string) {
// 			return i, s
// 		}
// 	}()
// 	// 执行returnFunc闭包
// 	ret1, ret2 := returnFunc(1, "Hello")
// 	// 打印闭包
// 	fmt.Printf("call returnFunc: %d, %s\n", ret1, ret2)
// 	fmt.Println("a.i = ", a.i)
// 	fmt.Println("after call function1, a.i = ", function1(1))
// 	fmt.Println("a.i = ", a.i)
// }

// func custom() {
// 	fmt.Println("Hello, World!")
// }

// func customHasReturn() int {
// 	return 2
// }
