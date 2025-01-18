package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//计算pai值

	var inCnt float64
	var outCnt float64
	for i := 0; i < 10000000; i++ {
		x := rand.Float32()
		y := rand.Float32()
		if x*x+y*y <= 1 {
			inCnt++
		}
		outCnt++
	}
	pi := 4 * inCnt / outCnt
	fmt.Println(pi)
}
