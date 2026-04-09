package main

import (
	"fmt"
)

func Add(a int, b int) {
	fmt.Println(a + b)
}

func Sub(a int, b int) {
	fmt.Println(a - b)
}

func AddChan(nums []int, c chan int) {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	c <- sum
}

func main() {

	// go Add(2, 3)

	// go Sub(3, 1)
	// fmt.Println("hello")
	// fmt.Println("raza")

	// time.Sleep(time.Second)

	nums := []int{1, 23, 4, 5, 6}
	c := make(chan int)

	go AddChan(nums[:len(nums)/2], c)
	go AddChan(nums[len(nums)/2:], c)

	x, y := <-c, <-c

	fmt.Println(x, y, x+y)

}
