package main

func main() {
	s := "bcca"
	t := "abc"
	println(validSubstringCount(s, t))
}

func validSubstringCount(s string, t string) int64 {
	if len(s) < len(t) {
		return 0
	}
	dif := [26]int{}
	for _, c := range t {
		dif[c-'a']++
	}
	cnt := 0
	for _, c := range dif {
		if c != 0 {
			cnt++
		}
	}
	pre := 0
	res := int64(0)
	for _, c := range s {
		dif[c-'a']--
		if dif[c-'a'] == 0 {
			cnt--
		}
		for cnt == 0 {
			if dif[s[pre]-'a'] == 0 {
				cnt++
			}
			dif[s[pre]-'a']++
		}
		res += int64(pre)
	}
	return res
}
