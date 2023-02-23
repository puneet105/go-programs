//program to check whether the sum of a number is 1 after adding square of each digit recursively
//happynumber

// A happy number is a number defined by the following process:

// Starting with any positive integer, replace the number by the sum of the squares of its digits.
// Repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1.
// Those numbers for which this process ends in 1 are happy.
// Return true if n is a happy number, and false if not.
package main

import "fmt"

func calculateSquare(number int) bool{
	sum := 0
	for number != 0 {
		x := number % 10
		sum = sum + (x * x)
		number = number/10
	}


	fmt.Println(sum)
	if sum == 1{
		//fmt.Println("sum is 1")
		return true
	}else{
		return calculateSquare(sum)
	}
	return false
}


func main() {
	number := 2

	check := calculateSquare(number)
	//fmt.Println(check)
	if check == true{
		fmt.Println("Happy Number")
	}

}

/*
//optimized way
package main

import (
	"fmt"
)
func isHappy(number int) bool{
	var slow, fast = number, number
	for{
		slow = calculateSquareNumber(slow)
		fmt.Println("Slow is: ", slow)
		fast = calculateSquareNumber(calculateSquareNumber(fast))
		fmt.Println("Fast is: ", fast)
		if slow == fast{
			break
		}
	}
	return slow == 1
}
func calculateSquareNumber(number int) int{
	sum := 0
	for number != 0{
		x := number%10
		sum = sum + (x * x)
		number = number/10
	}
	//fmt.Println(sum)
	return sum
}
func main(){
	fmt.Println("Hpappy Number Program")
	number := 2
	check := isHappy(number)
	fmt.Println("Check is: ", check)
	if check{
		fmt.Println("Happy Number It is...!!!")
	}else{
		fmt.Println("Not A Happy Number")
	}
}
*/