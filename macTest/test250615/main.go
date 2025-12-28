package main

import "fmt"

func main() {

}

func maxDiff(num int) int {
	baseArr := []int{}
	for num > 0 {
		temp := num % 10
		num /= 10
		baseArr = append(baseArr, temp)
	}
	iMax := 0
	baseNum := 9
	for i := len(baseArr) - 1; i >= 0; i-- {
		if baseNum == 9 && baseArr[i] != 9 {
			baseNum = baseArr[i]
		}
		if baseNum != 9 && baseArr[i] == baseNum {
			iMax = iMax*10 + 9
		} else {
			iMax = iMax*10 + baseArr[i]
		}
	}

	iMin := 0
	if baseArr[len(baseArr)-1] == 1 {
		baseNum = 0
		for i := len(baseArr) - 1; i >= 0; i-- {
			if baseNum == 0 && baseArr[i] != 0 && baseArr[i] != baseArr[len(baseArr)-1] {
				baseNum = baseArr[i]
			}
			if baseNum != 0 && baseArr[i] == baseNum {
				iMin = iMin*10 + 0
			} else {
				iMin = iMin*10 + baseArr[i]
			}
		}
	} else {
		iMin = 0
		baseNum = baseArr[len(baseArr)-1]
		for i := len(baseArr) - 1; i >= 0; i-- {
			if baseArr[i] == baseNum {
				iMin = iMin*10 + 1
			} else {
				iMin = iMin*10 + baseArr[i]
			}
		}
	}
	fmt.Println(iMax)
	fmt.Println(iMin)
	return iMax - iMin
}
