package main

import (
	"fmt"
	"math"
	"math/rand/v2"
	"runtime"
	"time"

	"rsc.io/quote"
)

func add(a, b int) (string, bool) {
	println(a + b)
	return "gulistane", true
}

func gretting() {
	fmt.Println("Hello, World!", "Raza")
	fmt.Println(quote.Go())

	fmt.Println("random number=", rand.Float32())

	fmt.Println("sqrt root of 144 is: ", math.Sqrt(142))

	fmt.Println("PI value is: ", math.Pi)

	// fmt.Println("Add to number", add(3, 7))
	c, d := add(4, 5)

	println(c, d)

	var name bool
	println(name)
}

func datatype() {
	a := 2.6
	// com := cmplx.Sqrt(5 + 2i)

	var x bool
	var y string
	var z int
	var w float32

	println("hello", x, y, z, w)

	b := int(a)
	println(b)

	const Pi = 3
	println(Pi)
}

func loop() {
	// sum := 1
	// for i := 1; i <= 5; i++ {
	// 	sum *= i
	// }

	// fmt.Println(sum)

	// isRun := true
	// b := 1
	// for isRun == true {
	// 	fmt.Println("raza")
	// 	if b == 5 {
	// 		break
	// 	}
	// 	b++

	// }

	// var num string = "raza"

	// if num == "Raza" {
	// 	name := "Gulistane"
	// 	println("true", name)
	// } else {
	// 	fmt.Println("false-")
	// }

	a := 1
	switch a {
	case 1:
		fmt.Println("monday")
	case 2:
		fmt.Println("monday")
	default:
		fmt.Println("default case run")
	}

	os := runtime.GOOS
	defer fmt.Println(os)

	fmt.Println(time.Now().Weekday())

}

func pointer() {

	a := 2

	b := &a

	a = 4

	fmt.Println("b value is = ", *b, a)

}

type Vertex struct {
	X int
	Y int
}

type ModerationItem struct {
	Type       string    `json:"type"` // "post" or "comment"
	ID         int       `json:"id"`
	PostID     int       `json:"postId,omitempty"`
	PostNumber int       `json:"postNumber,omitempty"`
	PostSlug   string    `json:"postSlug,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"createdAt"`
	PostTitle  string    `json:"postTitle,omitempty"`
}

type Person struct {
	ID        int
	FirstName string
	LastName  string
	RollNo    int
}

func main() {
	// defer greting()

	// datatype()

	// loop()

	// pointer()

	// fmt.Println(Vertex{2, 4})

	person1 := Person{1, "Gulistane", "Raza", 101}
	fmt.Println(person1.FirstName)
	fmt.Println(person1.LastName)

}
