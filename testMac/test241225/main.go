package main

import "fmt"

func main() {
	s1 := "11"
	s2 := "222222222aaaaaaaa"
	s3 := s1 + s2

	fmt.Println(&s1)
	fmt.Println(&s2)
	fmt.Println(&s3)
}
