//multi part upload logic using go routines to make it efficient
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func upload(file int, wg *sync.WaitGroup, uploadTime int){
	defer wg.Done()
	time.Sleep(time.Duration(uploadTime) * time.Second)
	fmt.Println("File ", file, " uploaded successfully and it took ", time.Duration(uploadTime) * time.Second, " seconds")
}
func main (){
	fmt.Println("Upload File Program")
	wg := sync.WaitGroup{}
	for i:=1;i<=10;i++{
		wg.Add(1)
		rand.Seed(time.Now().UnixNano())
		uploadTime := rand.Intn(20)
		go upload(i, &wg, uploadTime)
	}
	wg.Wait()
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
}