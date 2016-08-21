package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13} //array is fixed 6 elements

	var s []int = primes[1:4] //slices are dynamically-sized
	fmt.Println(s)
}
