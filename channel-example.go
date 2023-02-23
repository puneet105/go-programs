//Go Program to print odd and even number using go routines and channel
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func oddFunc(ch chan int, wg *sync.WaitGroup) {
	runtime.Gosched()
	for x := range ch {
		fmt.Println("Odd Number is", x)
	}
}

func evenFunc(ch chan int, wg *sync.WaitGroup) {
	for x := range ch {
		fmt.Println("Even number is", x)
	}

	close(ch)
}

func main() {
	var numbers = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	oddChan := make(chan int)
	evenChan := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go oddFunc(oddChan, &wg)
	go evenFunc(evenChan, &wg)
	for _, val := range numbers {
		if val%2 != 0 {
			oddChan <- val
		} else {
			evenChan <- val
		}
	}
	close(oddChan)
	close(evenChan)
	wg.Done()
	wg.Done()
	wg.Wait()
}
