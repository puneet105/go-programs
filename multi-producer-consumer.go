//o program to solve multiple producer and single consumer
package main

import (
"fmt"
"sync"
)

var messages = [][]string{{"message 1","message 2"}, {"message 3","message 4"}, {"message 5","message 6"}, {"message 7","message 8"}, {"message 9","message 10"}}
var producer_count int = 5
func multiproducer(job chan string,count int, wg *sync.WaitGroup){
	defer wg.Done()
	for _,msg := range messages[count]{
		fmt.Printf("Producer %v is producing %v message",count+1,msg)
		fmt.Println()
		job<-msg
	}
}

func multiconsumer(job  chan string, done chan <-bool){
	for msg := range job{
		fmt.Println("Message consumed is: ",msg)
	}
	done<-true
}

func main(){
	var wg sync.WaitGroup
	job := make(chan string)
	done := make(chan bool)
	for i:=0;i<producer_count;i++{
		wg.Add(1)
		go multiproducer(job,i,&wg)
	}
	go multiconsumer(job, done)

	wg.Wait()
	close(job)
	<-done

}
