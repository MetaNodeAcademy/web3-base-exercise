package exercise

import (
	"fmt"
	"time"
)

// 循环控制
func LoopControl() {
	fmt.Println("------Loop Control------")

	// 方式1
	// for i := 0; i < 10; i++ {
	// 	fmt.Printf("方式1 第 %d 次循环\n", i)
	// }
	// 方式2
	// b := 1
	// for b <= 10 {
	// 	fmt.Printf("方式2 第 %d 次循环\n", b)
	// 	b++
	// }
	// 方式3
	// ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second*2))
	// var started bool
	// var stopped atomic.Bool
	// for {
	// 	if !started {
	// 		fmt.Println("方式3 启动")
	// 		started = true
	// 		go func() {
	// 			for {
	// 				fmt.Println("方式3 运行中")
	// 				select {
	// 				case <-ctx.Done():
	// 					fmt.Println("select 停止")
	// 					stopped.Store(true)
	// 				default:
	// 					fmt.Println("select default 运行中")
	// 				}
	// 			}
	// 		}()
	// 	}
	// 	fmt.Println("方式3 循环")
	// 	if stopped.Load() {
	// 		fmt.Println("方式3 停止")
	// 		break
	// 	}
	// }
	// 方式4 遍历数组
	// var a [10]string
	// a[0] = "Hello"
	// for i := range a {
	// 	fmt.Printf("当前下标 %d\n", i)
	// }
	// 遍历切片
	// s := make([]string, 10)
	// s[0] = "Hello"
	// for i := range s {
	// 	fmt.Printf("当前下标 %d\n", i)
	// }
	// for i, v := range s {
	// 	fmt.Printf("当前下标 %d, 值 %s\n", i, v)
	// }

	// fmt.Println("------遍历map------")
	m := make(map[string]string)
	m["b"] = "Hello, b"
	m["a"] = "Hello, a"
	m["c"] = "Hello, c"
	for i := range m {
		fmt.Printf("当前key %s\n", i)
	}
	for k, v := range m {
		fmt.Printf("m[%s] = %s\n", k, v)
	}
}

// 1.9.4 循环控制语句
func LoopBreakControl() {
	fmt.Println("------循环控制语句------")
	// 中断for循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("break i == 5")
			break
		}
		fmt.Printf("第%d次循环\n", i)
	}
	// 中断switch
	fmt.Println("------中断switch------")
	switch i := 1; i {
	case 1:
		fmt.Println("i == 1")
		if i == 1 {
			break
		}
		fmt.Println("i == 1")
	case 2:
		fmt.Println("i == 2")
	case 3:
		fmt.Println("i == 3")
	default:
		fmt.Println("default")
	}
	// 中断select
	fmt.Println("------中断select------")
	select {
	case <-time.After(time.Second * 2):
		fmt.Println("过了2秒")
	case <-time.After(time.Second):
		fmt.Println("过了1秒")
		if true {
			break
		}
		fmt.Println("break 之后")
	}
	// 不使用标记
	fmt.Println("------不使用标记------")
	for i := 1; i <= 3; i++ {
		fmt.Printf("不使用标记,外部循环, i = %d\n", i)
		for j := 5; j <= 10; j++ {
			fmt.Printf("不使用标记,内部循环, j = %d\n", j)
			break
		}
	}
	// 使用标记
	fmt.Println("------使用标记------")
outter:
	for i := 1; i <= 3; i++ {
		fmt.Printf("使用标记,外部循环, i = %d\n", i)
		for j := 5; j <= 10; j++ {
			fmt.Printf("使用标记,内部循环, j = %d\n", j)
			break outter
		}
	}
}

// 1.9.4.2 continue语句
func ContinueControl() {
	fmt.Println("------continue语句------")
	for i := 0; i < 5; i++ {
		if i == 3 {
			fmt.Printf("continue %d\n", i)
			continue
		}
		fmt.Printf("第 %d 轮循环\n", i)
	}
	// 不使用标记
	fmt.Println("------continue语句,不使用标记------")
	for i := 1; i <= 2; i++ {
		fmt.Printf("不使用标记,外部循环, i = %d\n", i)
		for j := 5; j <= 10; j++ {
			fmt.Printf("不使用标记,内部循环, j = %d\n", j)
			if j >= 7 {
				fmt.Printf("不使用标记,内部循环, 跳出当前循环, j = %d\n", j)
				continue
			}
			fmt.Println("不使用标记,内部循环, 跳出外部循环")
		}
	}
	// 使用标记
	fmt.Println("------使用标记------")
outter:
	for i := 1; i <= 3; i++ {
		fmt.Printf("使用标记,外部循环, i = %d\n", i)
		for j := 5; j <= 10; j++ {
			fmt.Printf("使用标记,内部循环, j = %d\n", j)
			if j >= 7 {
				continue outter
			}
			fmt.Println("不使用标记,内部循环, 在continue之后执行")
		}
	}
}

// 1.9.4.3 goto语句
func GotoControl() {
	fmt.Println("------goto语句------")
	gotoPreset := false

preset:
	a := 5

process:
	if a > 0 {
		a--
		fmt.Println("当前a的值为: ", a)
		goto process
	} else if a <= 0 {
		if !gotoPreset {
			gotoPreset = true
			goto preset
		} else {
			goto post
		}
	}

post:
	fmt.Println("main将结束, 当前a的值为: ", a)

}
