package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Sarturday?")

	today := time.Now().Weekday()

	switch time.Saturday {
	case today + 0:
		fmt.Println("Today .")
	case today + 1:
		fmt.Println("Tomorrow .")
	case today + 4:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away")
	}
}
