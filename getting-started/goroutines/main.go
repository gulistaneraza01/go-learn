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

	// nums := []int{1, 23, 4, 5, 6}
	// c := make(chan int)

	// go AddChan(nums[:len(nums)/2], c)
	// go AddChan(nums[len(nums)/2:], c)

	// x, y := <-c, <-c

	// fmt.Println(x, y, x+y)

	// ch := make(chan int, 3)

	// ch <- 22
	// ch <- 223
	// ch <- 112

	// fmt.Println(<-ch, <-ch, <-ch)

	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)

}

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
