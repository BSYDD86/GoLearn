package main

import (
	"fmt"
	"sort"
)

func main() {
	str := ""
	fmt.Println(checkCode(str))
}

type coupon struct {
	code string
	val  string
}

func validateCoupons(code []string, businessLine []string, isActive []bool) []string {
	res := make([]coupon, 4)
	for i := range code {
		if !checkCode(code[i]) {
			isActive[i] = false
		} else if !checkBusinessLine(businessLine[i]) {
			isActive[i] = false
		}
		if isActive[i] {
			res = append(res, coupon{code: code[i], val: businessLine[i]})
		}
	}
	sort.Slice(res, func(i, j int) bool {
		if !(res[i].val == res[j].val) {
			return res[i].val < res[j].val
		} else {
			return res[i].code < res[j].code
		}
	})
	ans := []string{}
	for _, c := range res {
		ans = append(ans, c.code)
	}
	return ans
}
func checkCode(str string) bool {
	if str == "" {
		return false
	}
	for _, c := range str {
		if !(('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') || ('0' <= c && c <= '9') || c == '_') {
			return false
		}
	}
	return true
}

func checkBusinessLine(str string) bool {
	if !((str == "electronics") || (str == "grocery") || (str == "pharmacy") || (str == "restaurant")) {
		return false
	}
	return true
}
