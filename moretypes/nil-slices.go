/**
The zero value of a slice is nil.
値のないスライスはnilと同義です

A nil slice has a length and capacity of 0 and has no underlying array.
長さも容量ももたないスライスは、配列ではありません
*/
package main

import "fmt"

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))

	if s == nil {
		fmt.Println("nil!")
	}
}
