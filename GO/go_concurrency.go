package main

import (
	"fmt"
	"time"
)

func test (s string) {
	for i:=0; i < 5; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 500)
	}
}

func testing (s []int, c chan int) {
	sum := 0
	for _, i := range s {
		sum += i
	}
	c <- sum
}

func getVals (c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}

func selectFunc (c, quit chan int) {
	index := 0
	for {
		select {
		case c <- index :
			index += 1
		case <- quit:
			fmt.Println("quit")
			return
		}
	}
}

func main()  {

	//			GO Routines
	// go test("Hello")
	// go test("World")

	//			GO Channels
	// s := []int {7, 2, 8, -9, 4, 0}
	// c := make(chan int)
	// go testing(s[:len(s)/2], c)
	// go testing(s[len(s)/2:], c)
	// x, y := <- c, <- c
	// fmt.Println(x, y)

	//			GO Range and Close
	// c := make(chan int, 10)
	// for i := range c {
	// 	fmt.Println(i)
	// }
	// go getVals(c)

	//			GO select
	// c := make(chan int)
	// quit := make(chan int)
	//
	// go func()  {
	// 	for i:= 0; i < 10; i++ {
	// 		fmt.Println(<-c)
	// 	}
	// 	quit <- 0
	// }()
	// selectFunc(c, quit)


	//			GO select-default
	tick := time.Tick(200 * time.Millisecond)
	boom := time.After(2000 * time.Millisecond)
	for {
		select {
		case <- tick:
			fmt.Println("tick..")
		case <- boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
