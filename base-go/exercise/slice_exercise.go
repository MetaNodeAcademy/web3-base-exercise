package exercise

import (
	"fmt"
)

// 1.13 切片
func SliceExercise() {
	fmt.Println("-------切片-------")
	// 方式1，声明并初始化一个空的切片
	var s1 []int = []int{}
	// 方式2，类型推导，并初始化一个空的切片
	var s2 = []int{}
	// 方式3，与方式2等价
	s3 := []int{}
	// 方式4，与方式1、2、3等价，可以在大括号中定义切片的初始元素
	s4 := []int{1, 2, 3, 4, 5}
	// 方式5，用make()函数创建切片，创建[]int类型的切片，指定切片初始长度为0
	s5 := make([]int, 0)
	// 方式6，用make()函数创建切片，创建[]int类型的切片，指定切片初始长度为2，指定容量参数4
	s6 := make([]int, 2, 4)
	// 方式7，引用一个数组，初始化切片
	arr := [5]int{6, 5, 4, 3, 2}
	// 从数组下标2开始，直到数组的最后一个元素
	s7 := arr[2:]
	// 从数组下标1开始, 直到数组下标3的元素，创建一个新的切片
	s8 := arr[1:3]
	// 从0到下标2的元素，创建一个新的切片
	s9 := arr[:2]
	fmt.Println(s1, s2, s3, s4, s5, s6, arr, s7, s8, s9)
	fmt.Println(s7)
	fmt.Println(s8)
	fmt.Println(s9)
	// [5]int{9, 8, 7, 3, 2}
	arr[0] = 9
	arr[1] = 8
	arr[2] = 7
	fmt.Println(s7)
	fmt.Println(s8)
	fmt.Println(s9)

}

// 1.13.2 使用切片
func UseSlice() {
	fmt.Println("-------使用切片-------")
	// s1 := []int{5, 4, 3, 2, 1}
	// // 下标访问切片
	// e1 := s1[0]
	// e2 := s1[1]
	// e3 := s1[2]
	// fmt.Println(s1, e1, e2, e3)
	// // 向指定位置赋值
	// s1[0] = 10
	// s1[1] = 9
	// s1[2] = 8
	// fmt.Println(s1)
	// // range迭代访问切片
	// for i, v := range s1 {
	// 	fmt.Printf("当前下标 %d, 值 %d\n", i, v)
	// }
	// for i := range s1 {
	// 	fmt.Printf("当前下标 %d, 值 %d\n", i, s1[i])
	// }
	// for i := 0; i < len(s1); i++ {
	// 	fmt.Printf("当前下标 %d, 值 %d\n", i, s1[i])
	// }

	// s3 := []int{}
	// fmt.Println("s3 = ", s3)
	// // append函数追加元素
	// s3 = append(s3)
	// fmt.Println(len(s3), cap(s3))
	// s3 = append(s3, 1)
	// fmt.Println(len(s3), cap(s3))
	// s3 = append(s3, 2, 3)
	// fmt.Println(len(s3), cap(s3))
	// fmt.Println("s3 = ", s3)

	// s4 := []int{1, 2, 4, 5}
	// s4 = append(s4[:2], append([]int{3}, s4[2:]...)...)
	// s4 = append([]int{3}, s4...)
	// fmt.Println("s4 = ", s4)

	s5 := []int{1, 2, 3, 5, 4}
	s5 = append(s5[:2], s5[3:]...)
	fmt.Println("s5 = ", s5)
	//s5 = s5[2:] // 截取下标2之后的所有元素
	//s5 = s5[:3] // 截取下标3之前的所有元素
	s5 = s5[1:3] // 截取下标1到下标3之前的元素，不包含下标3
	fmt.Println("s5 = ", s5)

}

// 1.13.2.3 复制切片
func CopySlice() {
	fmt.Println("-------拷贝切片-------")
	src1 := []int{1, 2, 3}
	dst1 := make([]int, 4, 5)

	src2 := []int{1, 2, 3, 4, 5}
	dst2 := make([]int, 3, 3)
	fmt.Printf("before copy, src1 = %v dst1 = %v\n", src1, dst1)
	fmt.Printf("before copy, src2 = %v dst2 = %v\n", src2, dst2)
	copy(dst1, src1)
	copy(dst2, src2)
	fmt.Printf("after copy, src1 = %v dst1 = %v\n", src1, dst1)
	fmt.Printf("after copy, src2 = %v dst2 = %v\n", src2, dst2)
}

// 1.13.3 切片原理
func SlicePrinciple() {
	fmt.Println("--------切片原理---------")
	s := make([]int, 3, 6)
	fmt.Println("s length:", len(s))
	fmt.Println("s capacity:", cap(s))
	fmt.Println("s:", s)
	s[1] = 2
	fmt.Println("set position 1, s = ", s)

	modifySlice(s)
	fmt.Println("after modifySlice, s = ", s)
}

func modifySlice(param []int) {
	param[0] = 1024
}

// 没有触发切片扩容
func SliceNoExpand() {
	fmt.Println("--------没有触发切片扩容--------")
	s := make([]int, 3, 6)
	fmt.Println("initial s = ", s)
	s[1] = 2
	fmt.Println("after set position 1, s = ", s)

	s2 := append(s, 4)
	fmt.Println("after append, s length:", len(s))
	fmt.Println("after append, s capacity:", cap(s))
	fmt.Println("after append, s2 length:", len(s2))
	fmt.Println("after append, s2 capacity:", cap(s2))
	fmt.Println("after append, s = ", s)
	fmt.Println("after append, s2 = ", s2)

	s[0] = 1024
	fmt.Println("after set position 0, s = ", s)
	fmt.Println("after set position 0, s2 = ", s2)

	appendInFunc(s)
	fmt.Println("after append in func, s = ", s)
	fmt.Println("after append in func, s2 = ", s2)

}

func appendInFunc(param []int) {
	param = append(param, 1022)
	fmt.Println("in func, param = ", param)
	param[2] = 512
	fmt.Println("set position 2 in func, param = ", param)
}

// 触发切片扩容
func SliceExpand() {
	fmt.Println("--------触发切片扩容---------")
	s := make([]int, 2, 2)
	fmt.Println("initial s = ", s)

	s2 := append(s, 4)
	fmt.Println("after append, s length:", len(s))
	fmt.Println("after append, s capacity:", cap(s))

	fmt.Println("after append, s2 length:", len(s2))
	fmt.Println("after append, s2 capacity:", cap(s2))
	fmt.Println("after append, s = ", s)
	fmt.Println("after append, s2 = ", s2)

	s[0] = 1024
	fmt.Println("after set position 0, s = ", s)
	fmt.Println("after set position 0, s2 = ", s2)

	appendInFuncExpand(s2)
	fmt.Println("after append in func, s2 = ", s2)
}

func appendInFuncExpand(param []int) {
	param1 := append(param, 511)
	param2 := append(param1, 512)
	param1[0] = 1
	fmt.Println("in func, param1 = ", param1)
	param2[2] = 500
	fmt.Println("set position 2 in func, param2 = ", param2)
}