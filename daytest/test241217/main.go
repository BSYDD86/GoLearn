package main

import "fmt"

func main() {

	var st []string
	fmt.Println(cap(st))
	fmt.Println(st == nil)
	s := " 1 23  aa "
	//去除空格
	sArr := []byte(s)

	fast := 0
	slow := 0
	for fast < len(s) {
		if sArr[fast] == ' ' {
			fast++
		} else {
			break
		}
	}

	fmt.Println(slow)
	fmt.Println(fast)

	for ; fast < len(sArr); fast++ {
		if fast > 0 && sArr[fast-1] == ' ' && sArr[fast] == ' ' {
			continue
		} else {
			sArr[slow] = sArr[fast]
			slow++
		}
	}

	temp := make([]byte, slow)
	copy(temp, sArr)
	fmt.Println(string(temp))

}

//
//func reverseStr(s string, k int) string {
//
//	for l := 0; l < len(s); l += 2 * k {
//		tempL := l
//		tempR := min(l+k-1, len(s)-1)
//		for tempL < tempR {
//			s[1], s[tempR] = s[tempR], s[tempL]
//		}
//	}
//	return s
//}
