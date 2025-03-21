package main

import "fmt"

func main() {
	arr := [4]int{0, 1, 2, 3}
	fmt.Println(arr)
	slice := arr[:3]
	fmt.Println(slice)
	slice = append(slice, 4)
	fmt.Println(arr)
	temp := append(slice, 5)
	fmt.Println("-----")
	fmt.Println(temp)
	temp[3] = 10
	fmt.Println(arr)
	fmt.Println(temp)

}
