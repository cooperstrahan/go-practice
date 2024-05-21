package main

import (
	"fmt"
	"time"
)

func main() {
	// go count()
	// time.Sleep(time.Millisecond * 2)
	// fmt.Println("Hello Multithread")
	// time.Sleep(time.Millisecond * 5)

	c := make(chan int)
	a := []int{8, 6, 7, 5, 3, 0, 9, -1}

	go printCount(c)

	for _, v := range a {
		c <- v
	}
	time.Sleep(time.Millisecond * 1)
	fmt.Println("End of main")
}
