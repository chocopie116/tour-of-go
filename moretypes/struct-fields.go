package main

import "fmt"

type Vertext struct {
	X int
	Y int
}

func main() {
	v := Vertext{2, 3}
	fmt.Println(v.X)
	v.X = 5
	fmt.Println(v.X)
}
