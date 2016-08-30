/**
The empty interface

The interface type that specifies zero methods is known as the empty interface:
メソッド定義をもたないインターフェースは、empty interfaceで知られる。

interface{}
An empty interface may hold values of any type. (Every type implements at least zero methods.)
空のインターフェースは、どんな型の値ももつことができます。(言い換えるとインターフェースは、メソッド0でもOK)


Empty interfaces are used by code that handles values of unknown type. For example, fmt.Print takes any number of arguments of type interface{}.
空のインターフェースは、不明な型を制御する場合に使われます。
たとえばfmt.Printは 引数の数や型の不明なものを引数として取り扱うことができるようにinterface{}を渡すようにしています

*/
package main

import "fmt"

func main() {
	var i interface{}
	describe(i)

}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
