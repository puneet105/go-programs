//Go program to solve producer consumer problem using go routines and channels
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var msg = []string{"aaaaa", "bbbbb", "ccccc", "ddddd"}

func producer(job chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, x := range msg {
		fmt.Println("Producer produces msg: ", x)
		job <- x
		runtime.Gosched()
	}
	close(job)
}

func consumer(job chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for x := range job {
		fmt.Println("Received data from producer: ", x)
	}

}
func main() {
	wg := sync.WaitGroup{}
	job := make(chan string)
	//done := make(chan bool)
	wg.Add(2)
	go producer(job, &wg)
	go consumer(job, &wg)

	wg.Wait()
}

/*
//without using wait group
package main


import (
	"fmt"
)

var msg = []string{"aaaaa","bbbbb","ccccc","ddddd"}


func producer(job chan <-string){
	for _,x := range msg{
		job <- x
	}
	close(job)
}

func consumer(job <-chan string, done chan<-bool){
	for x:= range job{
		fmt.Println("Received data from producer: ",x)
	}
	done<-true
}
func main(){
	job := make(chan string)
	done := make(chan bool)
	go producer(job)
	go consumer(job,done)
	<-done
}
*/

/*
package main

import (
	"encoding/json"
	"fmt"
	"sync"
)

type Employee struct{
	Name		string
	Contact		string
	Address 	string
	Age 		int
}

var Employees =  []Employee{
	{
		Name: 		"Puneet Sharma",
		Contact: 	"9677355991",
		Address: 	"Ajmer",
		Age: 		26,
	},
	{
		Name: 		"Mahesh Sharma",
		Contact: 	"9828279408",
		Address: 	"Ajmer",
		Age: 		53,
	},
	{
		Name: 		"Krishna Sharma",
		Contact: 	"9461383390",
		Address: 	"Ajmer",
		Age: 		51,
	},
	{
		Name: 		"Jyoti Sharma",
		Contact: 	"8058033963",
		Address: 	"Ajmer",
		Age: 		25,
	},
}

func producer(job chan <- Employee, wg *sync.WaitGroup, done chan bool){
	defer wg.Done()
	for _,val:= range Employees{
		bytes, err := json.MarshalIndent(&val, ""," ")
		if err != nil{
			fmt.Println(err)
		}
		fmt.Println("Producer has done producing employee: ",string(bytes))
		job <- val
		done <- true
	}
	close(job)
}

func consumer(job <- chan Employee, wg *sync.WaitGroup, done chan bool){
	defer wg.Done()
	for val := range job{
		bytes, err := json.MarshalIndent(&val, ""," ")
		if err != nil{
			fmt.Println(err)
		}
		fmt.Println("Consumer has done consuming employee: ", string(bytes))
		<-done
	}
}
func main(){
	wg := sync.WaitGroup{}
	wg.Add(2)
	job := make(chan Employee)
	done := make(chan bool)
	go producer(job, &wg ,done)
	go consumer(job, &wg, done)
	wg.Wait()
}
*/