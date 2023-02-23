package main

import (
	"fmt"
	"runtime"
	"sync"
)

func printCapitalLetter(wg *sync.WaitGroup, done chan <- bool) {
	defer wg.Done()
	for c := 'A'; c <= 'Z'; c = c + 1 {
		fmt.Print(string(c))
		//fmt.Println()
		done<-true
	}
	close(done)
}

func printSmallLetter(wg *sync.WaitGroup, done <- chan bool) {
	defer wg.Done()
	for i := 'a'; i <= 'z'; i=i+1 {
		<-done
		fmt.Print(string(i))
		//fmt.Println( )
	}
}

func main() {
	var wg sync.WaitGroup
	done := make(chan bool)
	wg.Add(2)
	go printSmallLetter(&wg, done)
	go printCapitalLetter(&wg, done)
	fmt.Println("Number of go routines after: ", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println()
	fmt.Println("Number of go routines after: ", runtime.NumGoroutine())
}
