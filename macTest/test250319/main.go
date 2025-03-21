package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	var wg sync.WaitGroup
	var x chan int
	x = make(chan int)

	// 启动一个 goroutine 来处理通道中的数据
	wg.Add(1)
	go func(a chan int) {
		defer wg.Done()
		for {
			// 检查通道是否关闭
			t, ok := <-a
			if !ok {
				break
			}
			fmt.Println(t)
		}
	}(x)

	// 向通道发送一些数据
	for i := 0; i < 5; i++ {
		x <- i
	}

	// 确保只关闭通道一次
	once.Do(func() {
		close(x)
	})

	// 等待所有 goroutine 完成
	wg.Wait()
}
