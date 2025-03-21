package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(math.MaxInt == math.MaxInt64)

}
func decodeString(s string) string {
	res := ""
	cnt := 0

	stNum := []int{}
	stStr := []string{}

	for _, c := range s {
		if c == '[' {
			stNum = append(stNum, cnt)
			cnt = 0
			stStr = append(stStr, res)
			res = ""
		} else if c == ']' {
			tempStr := stStr[len(stStr)-1]
			stStr = stStr[:len(stStr)-1]
			tempCnt := stNum[len(stNum)-1]
			stNum = stNum[:len(stNum)-1]
			res = tempStr + strings.Repeat(res, tempCnt)
		} else if '0' <= c && c <= '9' {
			num, _ := strconv.Atoi(string(c))
			cnt = cnt*10 + num
		} else {
			res += string(c)
		}
	}
	return res
}
