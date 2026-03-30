package main

import "fmt"

type Number struct {
	Num1 int
	Num2 int
}

func (n Number) add() int {
	return n.Num1 + n.Num2
}

func main() {
	addTwoNumber := Number{2, 3}

	result := addTwoNumber.add()

	fmt.Println(result)
}

func fibo(n int) int {
	if n <= 1 {
		return n
	}
	return fibo(n-1) + fibo(n-2)
}
