package main

import "fmt"

func fibonacci() func() int {
	n := 0
	return func() int {
		r := fib(n)
		n += 1
		return r
	}
}

func fib(n int) int {
	if n == 1 || n == 0 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

func main() {
	f := fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
