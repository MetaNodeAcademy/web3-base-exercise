package exercise

import (
	"fmt"
	"time"
)

// 1.15.1 range迭代
func RangeExercise() {
	fmt.Println("--------range迭代---------")
	str1 := "abc123"
	for index := range str1 {
		fmt.Printf("str1 -- index: %d, value: %c\n", index, str1[index])
	}

	str2 := "测试中文"
	for index := range str2 {
		fmt.Printf("str2 -- index: %d, value: %d\n", index, str2[index])
	}
	fmt.Printf("len(str2) = %d\n", len(str2))
	runesFromStr2 := []rune(str2)
	buytesFromStr2 := []byte(str2)
	fmt.Printf("len(runesFromStr2) = %d\n", len(runesFromStr2))
	fmt.Printf("len(buytesFromStr2) = %d\n", len(buytesFromStr2))

	for index, value := range str2 {
		fmt.Printf("str2 -- index: %d, value: %d\n", index, str2[index])
		fmt.Printf("str2 -- index: %d, value: %d\n", index, value)
	}
}

// 1.15.2 对数组与切片迭代
func ArrayAndSliceIteration() {
	fmt.Println("--------对数组与切片迭代---------")
	array := [...]int{1, 2, 3}
	slice := []int{4, 5, 6}
	// 方法1: 只拿到数组的下标索引
	for index := range array {
		fmt.Printf("array -- index=%d value=%d \n", index, array[index])
	}
	for index := range slice {
		fmt.Printf("slice -- index=%d value=%d \n", index, slice[index])
	}
	// 方法2: 同时拿到数组的下标索引和对应的值
	for index, value := range array {
		fmt.Printf("array -- index=%d value=%d \n", index, array[index])
		fmt.Printf("array -- index=%d value=%d \n", index, value)
	}
	for index, value := range slice {
		fmt.Printf("slice -- index=%d value=%d \n", index, slice[index])
		fmt.Printf("slice -- index=%d value=%d \n", index, value)
	}
}

// 1.15.3 对通道的迭代
func ChannelIteration() {
	fmt.Println("--------- 1.15.3 对通道的迭代 ----------")
	ch := make(chan int, 10)

	go addData(ch)
	for i := range ch {
		fmt.Println(i)
		time.Sleep(time.Second * 1)
	}
}

func addData(ch chan int) {
	size := cap(ch)
	for i := 0; i < size; i++ {
		ch <- i
		//time.Sleep(time.Second * 1)
	}
	close(ch)
}

// 1.15.4 对映射集合迭代
func MapIteration() {
	fmt.Println("--------- 1.15.4 对映射集合迭代 ----------")
	hash := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
	}
	for key := range hash {
		fmt.Printf("key = %s, value = %d\n", key, hash[key])
	}
	for key, value := range hash {
		fmt.Printf("key = %s, value = %d\n", key, value)
	}
}
