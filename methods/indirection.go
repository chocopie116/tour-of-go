package main

import "fmt"

type Vertex struct {
	X, Y float64
}

//vは変数でもポインタでもOK
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

//vはポインタのみ許可
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	//変数
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10) //ポインタを渡す

	//ポインタ
	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
