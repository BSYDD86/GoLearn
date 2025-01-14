package main

func main() {
	for i, j := 0, 3; i <= 2 && j <= 3; i, j = i+1, j-1 {
		println(i + j)
	}
}
