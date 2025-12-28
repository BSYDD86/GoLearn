package main

import (
	"fmt"
	"strings"
)

type MyQueue struct {
	lSt []int
	rSt []int
}

func Constructor() MyQueue {
	return MyQueue{[]int{}, []int{}}
}

func main() {
	str := "\"MyCircularDeque\",\"insertLast\",\"insertLast\",\"insertFront\",\"insertFront\",\"getRear\",\"isFull\",\"deleteLast\",\"insertFront\",\"getFront\""
	split := strings.Split(str, ",")
	fmt.Println(len(split))

	str1 := "[3],[1],[2],[3],[4],[],[],[],[4],[]"
	split1 := strings.Split(str1, ",")
	fmt.Println(len(split1))

}
