// Mutex and Race Condition Handling


package main

import (
    "fmt"
    "sync"
)

type SafeCounter struct {
    mu sync.Mutex
    val int
}


func (c *SafeCounter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.val++
}


func (c *SafeCounter) Get() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.val
}

func main() {
    var wg sync.WaitGroup
    counter := SafeCounter{}

    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            counter.Increment()
        }()
    }

    wg.Wait()
    fmt.Println("Final Counter Value:", counter.Get())
}