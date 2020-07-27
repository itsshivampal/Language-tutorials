package main

import "fmt"

type I interface {
	M()
}

func printString (i I) {
	i.M()
}

type T struct {
	S string
}

func (t T) M () {
	fmt.Println(t.S)
}

func main () {
	t1 := T {"Shivam Pal"}
	t1.M()
	
	printString(t1)

	var i1 I
	i1 = t1
	i1.M()
}
