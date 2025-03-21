package main

import "fmt"

const (
	beijing  = iota + 1
	shanghai = iota
)

func main() {
	const length = 10
	fmt.Println(beijing)

}
