//Go program to print letter and numbers in alternate order using go routines and channels
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func printLetter(wg *sync.WaitGroup, letter chan<- bool, number <-chan bool) {
	defer wg.Done()
	for c := 'A'; c <= 'Z'; c = c + 1 {
		letter <- true
		fmt.Print(string(c))
		<-number
	}
	close(letter)
}

func printNumber(wg *sync.WaitGroup, number chan<- bool, letter <-chan bool) {
	defer wg.Done()
	for i := 1; i <= 28; i++ {
		number <- true
		fmt.Print(i)
		_, ok := <-letter
		if ok == false {
			break
		}
	}
}

func main() {
	var wg sync.WaitGroup
	letter := make(chan bool)
	number := make(chan bool)
	wg.Add(2)
	go printLetter(&wg, letter, number)
	go printNumber(&wg, number, letter)
	fmt.Println("Number of go routines before: ", runtime.NumGoroutine())
	<-number
	wg.Wait()
	fmt.Println("\n")
	fmt.Println("Number of go routines after: ", runtime.NumGoroutine())
}
