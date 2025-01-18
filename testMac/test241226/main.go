package main

//
//import "fmt"
//
//func main() {
//	a := uint32(101)
//	b :=  a<<1和1<<a
//	c := a >> 1
//	fmt.Println(a)
//	fmt.Println(b)
//	fmt.Println(c)
//
//}

////excel convert
//func convert(n int)string  {
//	if n<0{
//		return "exception"
//	}
//	res:=""
//	temp:=
//}

//func main() {
//	arr := []int{1, 3, 4, 5, 2, 7, 1, 9, 1, 1, 1, 1}
//	quick(arr, 0, len(arr)-1)
//	fmt.Println(arr)
//
//}

//// []
//func quick(arr []int, l, r int) {
//	if l == r {
//		return
//	}
//	idx := partition(arr, l, r)
//	if idx > l {
//		quick(arr, l, idx-1)
//	}
//	if idx < r {
//		quick(arr, idx+1, r)
//	}
//}
//
//func partition(arr []int, l, r int) int {
//	pivotalIdx := rand.Intn(r-l+1) + l
//	arr[pivotalIdx], arr[r] = arr[r], arr[pivotalIdx]
//	pivotal := arr[r]
//	preIdx := l
//	for i := l; i < r; i++ {
//		if arr[i] <= pivotal {
//			arr[i], arr[preIdx] = arr[preIdx], arr[i]
//			preIdx++
//		}
//	}
//	arr[r], arr[preIdx] = arr[preIdx], arr[r]
//	return preIdx
//}
