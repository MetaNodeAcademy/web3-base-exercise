package exercise

import (
	"fmt"
	"sync"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * 100)
		fmt.Printf("i = %d s = %s\n", i, s)
	}
}

// 并发-goroutine 与 channel
func GoroutineChannel() {
	fmt.Println("-----并发-goroutine 与 channel-----")
	go func() {
		fmt.Println("run goroutine in closure")
	}()
	go func(s string) {
		fmt.Println(s)
	}("gorouine: closure params")
	go say("in goroutine: world")
	say("hello")
	time.Sleep(time.Millisecond * 3000)
	fmt.Println("main goroutine end")
}

// 线程安全的计数器
type SafeCounter struct {
	mu    sync.Mutex
	count int
}

// 增加计数
func (c *SafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

// 获取当前计数
func (c *SafeCounter) GetCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

type UnsafeCounter struct {
	count int
}

// 增加计数
func (c *UnsafeCounter) Increment() {
	c.count += 1
}

// 获取当前计数
func (c *UnsafeCounter) GetCount() int {
	return c.count
}

// Go 并发线程安全
func ConcurrentThreadSafe() {
	fmt.Println("-----Go 并发线程安全-----")
	safeCounter := SafeCounter{}
	// 启动100个goroutine同时增加计数
	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				safeCounter.Increment()
			}
		}()
	}

	unSafeCounter := UnsafeCounter{}
	// 启动100个goroutine同时增加计数
	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				unSafeCounter.Increment()
			}
		}()
	}
	// 等待一段时间确保所有goroutine完成
	time.Sleep(time.Second * 3)
	// 输出最终计数
	fmt.Printf("Final safeCounter: %d unSafeCounter:%d\n", safeCounter.GetCount(), unSafeCounter.GetCount())
}

// Channel
func ChannelExercise() {
	fmt.Println("-----Channel-----")
	// 创建一个带缓冲的channel
	ch := make(chan int, 3)
	// 启动发送goroutine
	go sendOnly(ch)
	// 启动接收goroutine
	go receiveOnly(ch)
	// 使用select进行多路复用
	timeout := time.After(time.Second * 10)
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				fmt.Println("channel closed")
				return
			}
			fmt.Printf("主goroutine接收到数据: %d\n", v)
		case <-timeout:
			fmt.Println("time out")
			return
		default:
			fmt.Println("没有数据, 等待中...")
			time.Sleep(time.Millisecond * 500)
		}
	}

}

// 只接收channel的函数
func receiveOnly(ch <-chan int) {
	for v := range ch {
		fmt.Printf("接收: %d\n", v)
		time.Sleep(time.Millisecond * 1000)
	}
}

// 只发送channel的函数
func sendOnly(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Printf("发送: %d\n", i)
		time.Sleep(time.Millisecond * 500)
	}
	close(ch)
}

// 并发 select
func GoroutineSelect() {
	fmt.Println("--------并发 select---------")
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	ch3 := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			// ch1 <- i
			// ch2 <- i
			// ch3 <- i
			// fmt.Printf("发送: %d\n", i)
			select {
			case ch1 <- i:
				fmt.Printf("ch1 发送: %d\n", i)
			case ch2 <- i:
				fmt.Printf("ch2 发送: %d\n", i)
			case ch3 <- i:
				fmt.Printf("ch3 发送: %d\n", i)
			}
		}
	}()
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch1:
			fmt.Printf("i = %d ch1 receive %d\n", i, x)
		case y := <-ch2:
			fmt.Printf("i = %d ch2 receive %d\n", i, y)
		case z := <-ch3:
			fmt.Printf("i = %d ch3 receive %d\n", i, z)
			// default:
			// 	fmt.Printf("receive nothing\n")
		}
	}

}
