package main

import (
	"fmt"
	"time"
)

func parTest(i int) {
	fmt.Println(i)

}

func main() {
	for i := 0; i < 10000; i++ {
		parTest(i)
		time.Sleep(100 * time.Millisecond)
	}
}
