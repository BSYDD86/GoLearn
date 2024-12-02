package main

import "fmt"

func main() {
	//slice := []int{1, 2, 3}
	//fmt.Printf("%d,%v", len(slice), slice)

	var slice2 []int
	slice2 = make([]int, 3)
	fmt.Printf("%d,%v", len(slice2), slice2)
	if slice2 == nil {
		fmt.Printf("yes")
	} else {
		fmt.Println("no")
	}

}
