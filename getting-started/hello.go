package main

import (
	"fmt"
	"math"
	"math/rand"

	"rsc.io/quote"
)

func add(a, b int) (string, bool) {
	println(a + b)
	return "gulistane", true
}

func main() {
	fmt.Println("Hello, World!", "Raza")
	fmt.Println(quote.Go())

	fmt.Println("random number=", rand.Float32())

	fmt.Println("sqrt root of 144 is: ", math.Sqrt(142))

	fmt.Println("PI value is: ", math.Pi)

	// fmt.Println("Add to number",add(3,7))
	c, d := add(4, 5)

	println(c, d)

	var name bool
	println(name)
}
