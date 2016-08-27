package main

import "fmt"

func main() {
	i, j := 42, 2701
	p := &i         //point to i
	fmt.Println(*p) //42

	fmt.Println(p) //pointer
	*p = 21        //set i through the pointer
	fmt.Println(i) //i = 21

	p = &j         //point to j
	*p = *p / 37   //devide j through pointer
	fmt.Println(j) //73
}
