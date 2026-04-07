package main

import "fmt"

type Number struct {
	Num1 int
	Num2 int
}

func (n Number) add() int {
	return n.Num1 + n.Num2
}

// type Vertex struct {
// 	X float64
// 	Y float64
// }

// func test(v *Vertex) {
// 	v.X = 100
// }

// func (v *Vertex) Scale(f float64) {
// 	v.X = v.X * f
// 	v.Y = v.Y * f
// }

// func ScaleFunc(v *Vertex, f float64) {
// 	v.X = v.X * f
// 	v.Y = v.Y * f
// }

// type Animal interface {
// 	Sound() string
// }

// type Dog struct {
// 	Name string
// 	Age  int
// }

// func (d Dog) Sound() string {
// 	return "Woof"
// }

func main() {
	// addTwoNumber := Number{2, 3}

	// result := addTwoNumber.add()

	// fmt.Println(result)

	// v := Vertex{1, 2}

	// test(&v)

	// fmt.Println(v)

	// v := Vertex{3, 4}
	// v.Scale(2)
	// ScaleFunc(&v, 10)

	// animal := Dog{"Buddy", 24}
	// fmt.Println(animal.Age, animal.Name)

	// var i interface{} = 1

	// value, err := i.(int)
	// fmt.Println(value, err)

	// var i interface{} = 4
	// do(i)

	_, err := Sqrt(4)
	fmt.Println(err)
	fmt.Println(Sqrt(-9))

}

type ErrNegativeSqrt int

func (e ErrNegativeSqrt) Error() string {
	// return fmt.PrintLn("cannot Sqrt negative number:", e)
	return "cannot Sqrt negative number"
}

func Sqrt(num int) (int, error) {

	if num < 0 {
		return 0, ErrNegativeSqrt(num)
	}

	return num * num, nil
}

// func random(i interface{}) {
// 	fmt.Println(i)
// }

// func do(i interface{}) {
// 	switch v := i.(type) {
// 	case int:
// 		fmt.Println("a", i, v)

// 	case string:
// 		fmt.Println("hello")

// 	default:
// 		fmt.Println("invalid")
// 	}
// }
