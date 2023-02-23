package main

import "fmt"

func printFibonacci(num int){
	x := 0
	y := 1
	fmt.Printf("Fibonacci Series is : %d %d", x,y)
	for i:=2;i<num;i++{
		z := x + y
		fmt.Printf(" %d", z)
		x = y
		y = z
	}
}

func main(){

	printFibonacci(5)
	fmt.Println()
	printFibonacci(10)
}