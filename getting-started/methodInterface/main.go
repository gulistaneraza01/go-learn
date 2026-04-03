package main

type Number struct {
	Num1 int
	Num2 int
}

func (n Number) add() int {
	return n.Num1 + n.Num2
}

type Vertex struct {
	X float64
	Y float64
}

func test(v *Vertex) {
	v.X = 100
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	// addTwoNumber := Number{2, 3}

	// result := addTwoNumber.add()

	// fmt.Println(result)

	// v := Vertex{1, 2}

	// test(&v)

	// fmt.Println(v)

	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

}
