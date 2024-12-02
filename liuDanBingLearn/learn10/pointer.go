package main

import "fmt"

func change(a *int, b *int) {
	var temp = *a
	*a = *b
	*b = temp
}

func main() {
	var a = 10
	var b = 100
	change(&a, &b)
	fmt.Println(a, b)
	fmt.Println(&a)
	var aidx = &a
	var aaidx **int
	aaidx = &aidx
	fmt.Println(aidx)
	fmt.Println(aaidx)

	fmt.Println(*aidx)
	fmt.Println(*&aaidx)
	fmt.Println(&*aaidx)

}
