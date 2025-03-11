// Find the missing number in array

package main

import (
    "fmt"
)


func findMissingHashSet(arr []int, n int) int{
    ns := make(map[int]bool)
    for _,num := range arr{
        ns[num] = true
    }
    fmt.Println(ns)

    for i:=1;i<n;i++{
        if !ns[i]{
            return i
        }
    }
    return -1
}

func main() {
    arr := []int{1, 2, 3, 4, 5, 6, 7, 9, 10}
    n := 10
    fmt.Println("Missing Number:", findMissingHashSet(arr, n))
}

