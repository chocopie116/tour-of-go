/**
A nil interface value holds neither value nor concrete type.
nilのインターフェースの値は値も、型ももちません。

Calling a method on a nil interface is a run-time error because there is no type inside the interface tuple to indicate which concrete method to call.
nilのインターフェースをよぶと、実行時エラーがおきます。
*/

package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I
	describe(i)
	i.M() //panic(runtime error)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
