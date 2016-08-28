package main

/**
Interface values with nil underlying values
インターフェースは、基底の値のnil値を評価します

If the concrete value inside the interface itself is nil, the method will be called with a nil receiver.
インターフェースを実装した変数・ポインタがの値がnilの場合、手続きはnilレシーバで手続きが呼び出されます


In some languages this would trigger a null pointer exception, but in Go it is common to write methods that gracefully handle being called with a nil receiver (as with the method M in this example.)
いくつかの言語では、こういった場合はnull pointer exceptionを引き起こします。
しかしGo言語では、nilレシーバとして呼び出されるものが一般的とされています。
サンプルにあるmethod M()がそれです。


Note that an interface value that holds a nil concrete value is itself non-nil.
インターフェースの値はnilではないものは、具体的な値を保持するものです
*/

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}

	fmt.Println(t.S)
}

func main() {
	var i I
	var t *T

	i = t
	describe(i)
	i.M()

	i = &T{"Hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
