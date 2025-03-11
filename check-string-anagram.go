// Check if strings are anagrams

package main

import (
    "fmt"
    "sort"
    // "strconv"
    "strings"
)

func sortStrings(x string) string{
    rune1 := []rune(x)
    sort.Slice(rune1, func(i, j int) bool {
        return rune1[i] < rune1[j]
    })

    return string(rune1)
}

func main (){
    str1 := "hello"
    str2 := "world"
    
    if len(str1) != len(str2){
        fmt.Println("Strings are not anagrams")
    }else{
        str1 = strings.ToLower(str1)
        str2 = strings.ToLower(str2)
        if sortStrings(str1) == sortStrings(str2){
            fmt.Println("Strings are anagrams")
        }else{
            fmt.Println("Strings are not anagrams")
        }
        
    }
   
}

//ANOTHER WAY

package main

import (
    "fmt"
)

func main(){
    s1 := "worth"
    s2 := "throw"

    ls1 := len(s1)
    ls2 := len(s2)

    am := make(map[string]int)
    for i:=0;i<ls1;i++{
        am[string(s1[i])]++
    }

    fmt.Println(am)
    for i:=0;i<ls2;i++{
        am[string(s2[i])]--
    }

    fmt.Println(am)
    for i:=0;i<ls1;i++{
        if am[string(s1[i])] != 0{
            fmt.Println("Not anagram")
            break
        }else{
            fmt.Println("anagram")
            break
        }
    }
    
    
}