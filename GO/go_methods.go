// running a go script - go run file_name.go
// formatting a go script - gofmt -w file_name.go

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	x, y float64
}

func (v Vertex) Abs () float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func (v *Vertex) Scale (f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func ScaleFunc (v *Vertex, f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func main () {
	v := Vertex {3, 4}
	v.Scale(2)
	fmt.Println(v.Abs())
	ScaleFunc(&v, 10)
	fmt.Println(v.Abs())

	p := &Vertex {4, 3}
	p.Scale(2)
	fmt.Println(p.Abs())
	ScaleFunc(p, 10)
	fmt.Println(p.Abs())

	fmt.Println(v, p)
}
