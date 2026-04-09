package main

import "fmt"

func Add[T comparable](num []T, a T) {
	fmt.Printf("%T, %T\n", num, a)
}

func main() {
	nums := []string{"1", "23", "4"}
	Add(nums, "6")
	fmt.Println("hello")
}
