// Q1: Implement a worker pool in Go where multiple workers process tasks concurrently. Requirements: Use goroutines and channels. Limit the number of concurrent workers. Ensure all tasks are processed efficiently.

package main

import (
    "fmt"
    "sync"
    "time"
)

type Task struct{
    ID int
}

func worker(i int, tasks <- chan Task, wg *sync.WaitGroup){
    defer wg.Done()
    for task := range tasks{
        fmt.Printf("Worker %d processed task %d\n",i,task.ID)
        time.Sleep(1*time.Second)
    }
    
}

func workerpool(wc int, tc int){
    tasks := make(chan Task, tc)
    var wg sync.WaitGroup

    for i:=1;i<=wc;i++{
        wg.Add(1)
        go worker(i,tasks,&wg)
    }

    for i:=1;i<=tc;i++{
        tasks <- Task{ ID: i}
    }

    close(tasks)

    wg.Wait()
}

func main(){
    wc := 4
    tc := 10

    workerpool(wc,tc)
}