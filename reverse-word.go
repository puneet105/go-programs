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

