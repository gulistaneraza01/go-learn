package main

import (
	"fmt"
	"math"
	"math/rand"

	"rsc.io/quote"
)

func main() {
    fmt.Println("Hello, World!", "Raza")
    fmt.Println(quote.Go())

    fmt.Println("random number=", rand.Float32())

	fmt.Println("sqrt root of 144 is: ",math.Sqrt(142))

	fmt.Println("PI value is: ",math.Pi)
}