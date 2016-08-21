package main

import (
	"fmt"
)

//newton法で平方根の近似値を出す
func SqrtNewton(x float64) float64 {
	x2 := x - (x*x-2)/(x*2) //Newton法

	return x2
}

//xのそれっぽい平方根を返す
func Sqrt(x float64) float64 {
	x2 := SqrtNewton(x)

	if x-x2 > 0.00001 {
		return Sqrt(x2)
	}

	return x2
}

func main() {
	fmt.Println("-----10 times ----")
	//10回Sqrt計算
	num := float64(2)
	for i := 0; i < 10; i++ {
		num = SqrtNewton(num)
		fmt.Println(num)
	}

	fmt.Println("----- loop ----")
	//値がいい感じになるまでSqrt
	fmt.Println(Sqrt(2))
}
