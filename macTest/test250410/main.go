package main

import (
	"fmt"
	"time"
)

func main() {
	n := 100000
	//fmt.Println(n)
	startTime := time.Now()
	fmt.Println(startTime)
	for i := 1; i <= n; i++ {
		for j := i; j <= n; j++ {
			//fmt.Println("i is", i, "j is", j)
		}
	}
	endTime := time.Now()
	fmt.Println(endTime)
	fmt.Println(endTime.Sub(startTime))
}
