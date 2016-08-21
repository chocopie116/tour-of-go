/**

A slice has both a length and a capacity.
スライス型は、長さと容量の2つを持ちます。

The length of a slice is the number of elements it contains.
スライスの長さは、保持している要素数を表します。

The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
スライスの容量とは基本的な配列の要素の数で、スライスの最初の要素から数えたものです

The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s).
スライスの長さと容量は、len(s)とcap(s)によってそれぞれ求めることができます。

You can extend a slice's length by re-slicing it, provided it has sufficient capacity. Try changing one of the slice operations in the example program to extend it beyond its capacity and see what happens.
スライスを再度スライスすると長さを継承することができ、十分な容量をもつものを提供できます。
容量を超えて継承をしようとスライスに操作を加えようとするとどうなるか確認しましょう。
*/

package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	//Slice the slice to give it zero length
	s = s[:0]
	printSlice(s)

	//Extend its length
	s = s[:4]
	printSlice(s)

	//Drop its first two values
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
