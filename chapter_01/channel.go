package main

import "fmt"

func printCount(c chan int) {
	num := 0
	for num >= 0 {
		fmt.Print(num, " ")
		num = <-c
		fmt.Print(num, " \n")
	}
}
