// running a go script - go run file_name.go
// formatting a go script - gofmt -w file_name.go


package main

import (
	"fmt"
	"math/rand"
	// "time"
)

var globalVar string = "This is global string"


func intSeq() func () int {
	i := 0
	return func () int {
		i += 1
		return i
	}
}


func plusNums(a int, b int, c string) int {
	fmt.Println(c, a + b)
	return 1
}

func pluPlus (a, b int, c string) int {
	fmt.Println(c, a + b)
	return 1
}


type person struct {
	name string
	age int
}


func main() {
	fmt.Println("Hello World!!")

	// --------------------------- Variable declaration ------------------------------//

	var myInt int = 10			// declaring variable in Go
	myInt2 := 12				// auto type deduction using :=
	fmt.Println(myInt + myInt2)

	var myString string = "Shivam"		// sting variable declaration
	myString2 := "Shivam Pal"			// auto type deduction
	fmt.Println(myString2 + " " + myString)

	fmt.Println(globalVar)		// global variable use

	// --------------------------- Array Operations ------------------------------//

	var arr [4]int				// declaring an array of size 4
	arr[0] = 1
	arr[2] = 2
	fmt.Println(arr)
	fmt.Println(arr[:2])

	var b1 [4]string; b1[1] = "Shivam"; b1[3] = "Pal"
	fmt.Println(b1)

	b:= [2]string{"test1", "test1"}
	fmt.Println(b)

	// --------------------------- Maps Operations ------------------------------//

	a := make(map[string]string)
	a["firstName"] = "Shivam"
	a["lastName"] = "Pal"
	fmt.Println(a)
	fmt.Println(a["firstName"])

	// --------------------------- If-else ------------------------------//

	num := rand.Intn(100)
	if num < 30 || num > 70 {
		fmt.Println(num, " You loose!!")
	} else {
		fmt.Println(num, " You win!!")
	}

	// --------------------------- If-else ------------------------------//

	num = rand.Intn(5)
	switch num {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case  3:
		fmt.Println("Three")
	default:
		fmt.Println("Another Value")
	}

	// --------------------------- For Loops ------------------------------//

	// 'for' is GO's only loooping construct

	i := 1
	for i <= 3 {
		fmt.Print(i, " ")
		i += 1
	}

	for j:= 1; j <= 3; j++ {
		fmt.Print(j, " ")
	}

	for {
		fmt.Println("Loop Finish!!")
		break
	}

	for j:= 0; j <= 9; j++ {
		if j%2 == 0 {
			fmt.Println("Even number!!")
		}
	}

	// --------------------------- Functions ------------------------------//

	plusNums(1, 2, "Sum of two nums ")
	pluPlus(8, 10, "Sum of two nums ")

	// --------------------------- CLosures ------------------------------//

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())


	// --------------------------- Structures ------------------------------//

	p1 := person{name: "Shivam", age: 24}
	p2 := person{name: "Somveer"}
	fmt.Println(p1.name, p1.age)
	fmt.Println(p2.name, p2.age)


}
