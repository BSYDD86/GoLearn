package main

//
//import "fmt"
//
////数轴上有n个点，当前你在第k个点。有一个指令字符串s，包含'L', 'R', '?'。L代表向左走，R代表向右走，?未知指令。现在问将?号替换成L或着R，对于数轴上的每一个点，是否可以访问到？输入描述第一行输入n和k
////第二行输入指令字符串s
////1 <= n <= 2e5, 1 <= k <= n
////s只由L,R,?组成
////1 <= len(s) <= 2e5输出描述输出一个长度为n的字符串ans，如果ans[i]为1，则可以访问到，为0则不可以访问到。
//
//////input
////5 2
////LL?
////
////output
////11000
//
//func main() {
//	arr := []int{1, 2}
//	fmt.Println(string(arr))
//}
//
//func confusingNumber(n int) bool {
//	res := []int{}
//	for n > 0 {
//		res = append(res, n%10)
//		n /= 10
//	}
//	r := len(res) - 1
//	l := 0
//	//[]
//	for l < r {
//		lconv := convert(res[l])
//		rconv := convert(res[r])
//		if lconv == -1 || rconv == -1 {
//			return false
//		}
//		if lconv != res[r] || rconv != res[l] {
//			return false
//		}
//		l++
//		r--
//	}
//	return true
//}
//func convert(n int) int {
//	if n == 0 || n == 1 || n == 6 || n == 8 || n == 9 {
//		if n == 6 {
//			return 9
//		}
//		if n == 9 {
//			return 6
//		}
//		return n
//	}
//	return -1
//}
