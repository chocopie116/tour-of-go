package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
<<<<<<< HEAD

=======
>>>>>>> 5877f8e1424f7d0a0b63dbd70817ba12e16049d6
func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
}
