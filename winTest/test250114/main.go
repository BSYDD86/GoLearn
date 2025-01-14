package main

func main() {

}
func lemonadeChange(bills []int) bool {
	five := 0
	ten := 0
	for bill := range bills {
		if bill == 5 {
			five++
		} else if bill == 10 {
			if five < 0 {
				return false
			} else {
				five--
				ten++
			}
		} else if bill == 20 {
			if ten >= 2 {
				ten -= 2
			} else if ten >= 1 && five >= 1 {
				ten--
				five--
			} else if five >= 3 {
				five -= 3
			} else {
				return false
			}
		}
	}
	return true
}
