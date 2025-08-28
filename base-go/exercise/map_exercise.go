package exercise

import (
	"fmt"
	"sync"
	"time"
)

// 1.14 Map集合
func MapExercise() {
	fmt.Println("-----Map集合-----")
	var m1 map[string]string
	fmt.Println("m1 = ", m1)

	m2 := make(map[string]string)
	fmt.Println("m2 length = ", len(m2))
	fmt.Println("m2 = ", m2)

	m3 := make(map[string]string, 10)
	fmt.Println("m3 length = ", len(m3))
	fmt.Println("m3 = ", m3)

	m4 := map[string]string{}
	fmt.Println("m4 length = ", len(m4))
	fmt.Println("m4 = ", m4)

	m5 := map[string]string{
		"name": "tom",
		"age":  "18",
	}
	fmt.Println("m5 length = ", len(m5))
	fmt.Println("m5 = ", m5)
}

// 1.14.2 使用Map集合
func UseMap() {
	fmt.Println("--------使用Map集合---------")
	m := make(map[string]int, 10)

	m["1"] = int(1)
	m["2"] = int(2)
	m["3"] = int(3)
	m["4"] = int(4)
	m["5"] = int(5)
	m["6"] = int(6)

	// 获取元素
	value1 := m["1"]
	fmt.Println("m[\"1\"] = ", value1)

	value1, exist := m["1"]
	fmt.Printf("m[\"1\"] = %d, exist = %t\n", value1, exist)

	valueUnexist, exist := m["10"]
	fmt.Printf("m[\"10\"] = %d, exist = %t\n", valueUnexist, exist)

	// 修改值
	fmt.Printf("before modify, m[\"2\"] = %d\n", m["2"])
	m["2"] = 20
	fmt.Printf("after modify, m[\"2\"] = %d\n", m["2"])

	// 修改map的长度
	fmt.Printf("before add, len(m) = %d\n", len(m))
	m["10"] = 10
	fmt.Printf("after add, len(m) = %d\n", len(m))

	// 遍历map集合
	for key, value := range m {
		fmt.Printf("iterate map, m[%s] = %d\n", key, value)
	}
	// 使用内置函数删除指定的key
	_, exist_10 := m["10"]
	fmt.Printf("before delete, exist 10: %t\n", exist_10)
	delete(m, "10")
	_, exist_10 = m["10"]
	fmt.Printf("after delete, exist 10: %t\n", exist_10)

	// 在遍历时，删除map中的key
	for key := range m {
		fmt.Printf("iterate map, delete key: %s\n", key)
		delete(m, key)
	}
	fmt.Printf("after delete, m: %v\n", m)
}

// 1.14.3 map作为参数
func MapAsParam() {
	fmt.Println("--------map作为参数-------")
	m := make(map[string]int)
	m["a"] = 1
	receiveMap(m)
	fmt.Printf("m = %v\n", m)
}

func receiveMap(param map[string]int) {
	fmt.Println("before modify, in receiveMap func: param[\"a\"] = ", param["a"])
	param["a"] = 2
	param["b"] = 3
}

// 1.14.4 并发时使用map集合
func ConcurrentUseMap() {
	fmt.Println("--------并发时使用map集合-------")
	m := make(map[string]int)
	var lock sync.Mutex
	go func() {
		for {
			lock.Lock()
			m["a"]++
			lock.Unlock()
		}
	}()
	go func() {
		for {
			lock.Lock()
			fmt.Printf("m[\"a\"] = %d\n", m["a"])
			lock.Unlock()
		}
	}()

	select {
	case <-time.After(time.Second * 5):
		fmt.Println("time out, stopping")
	}
}
