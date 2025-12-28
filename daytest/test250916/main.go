package main

func main() {

}

func canMeasureWater(x int, y int, target int) bool {
	temp := getGcd(x, y)
	return target%temp == 0
}

func getGcd(x, y int) int {
	a := max(x, y)
	b := min(x, y)
	for b != 0 {
		temp := b
		b = a % b
		a = temp
	}
	return a
}
