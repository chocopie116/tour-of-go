package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//パターンA
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

//パターンB
//func (v Vertex) Scale(f float64) {
//	v.X = v.X * f
//	v.Y = v.Y * f
//}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	//出力結果
	//パターンA: 50
	//パターンB: 5
	fmt.Println(v.Abs())
}
