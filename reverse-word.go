package main

import (
	"fmt"
	"strings"
)

func main(){
	str := "PUNEET SHARMA IS A SOFTWARE DEVELOPER"
	words := strings.Fields(str)

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	strings.Join(words, " ")
	fmt.Println(words)
}

//ANOTHER WAY

package main

import (
    "fmt"
    // "strconv"
)

func main (){
    str := "PUNEETSHARMA"
    lenght := len(str)
    final := []rune(str)
    fmt.Println("lenght is : ", lenght)
    for i:=0; i<lenght/2; i++{
        final[i], final[lenght-1-i] = final[lenght-1-i], final[i]
        
    }

    fmt.Println(string(final))
}

