package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(5, 12))      //(5^2 + 12^2)^1/2
	fmt.Println(compute(hypot))    //(3*3 + 4*4)^1/2
	fmt.Println(compute(math.Pow)) //3^4
}
