package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	getToday()
}

func getToday() {
	weekday := time.Now().Weekday()
	fmt.Println(weekday)
}

func getOs() {
	goos := runtime.GOOS
	fmt.Println(goos)
	switch goos {
	case "darwin":
		fmt.Println("macOS.")
	default:
		fmt.Println("unkown")

	}
}
