package main

import (
	"strconv"
	"strings"
)

const mx = 26

func main() {

}

func decodeString(s string) string {
	numSt := []int{}
	strSt := []string{}
	res := ""
	cnt := 0
	for _, c := range s {
		if c == '[' {
			numSt = append(numSt, cnt)
			strSt = append(strSt, res)
			res = ""
			cnt = 0
		} else if c == ']' {
			cnt = numSt[len(numSt)-1]
			numSt = numSt[:len(numSt)-1]
			temp := strSt[len(strSt)-1]
			strSt = strSt[:len(strSt)-1]
			res = temp + strings.Repeat(res, cnt)
			cnt = 0
		} else if '0' <= c && c <= '9' {
			temp, _ := strconv.Atoi(string(c))
			cnt = cnt*10 + temp
			if res != "" {
				strSt = append(strSt, res)
				res = ""
			}
		} else {
			res += string(c)
		}
	}
	for len(strSt) > 0 {
		temp := strSt[len(strSt)-1]
		strSt = strSt[:len(strSt)-1]
		res = temp + res
	}
	return res

}
