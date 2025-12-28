package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "-42"
	fmt.Println(myAtoi(str))
	strings.TrimSpace(str)
}

func myAtoi(str string) int {
	res := 0
	i := 0
	isNegative := false
	for i < len(str) {
		c := str[i]
		if isValid(c) || c == '-' {
			if c == '-' {
				isNegative = true
			} else {
				res = res*10 + int(c-'0')
			}
		}
		i++
	}
	if isNegative {
		res *= -1
	}
	return res
}
func isValid(c byte) bool {
	temp := int(c - '0')
	if temp >= 0 && temp <= 9 {
		return true
	}
	return false
}
