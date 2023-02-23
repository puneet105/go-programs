//character count in string
package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "Naveen"
	lower := strings.ToLower(name)

	//fmt.Println(lower)
	countChar := make(map[string]int, 0)
	for _, x := range lower {
		_, ok := countChar[string(x)]
		if ok {
			countChar[string(x)] = countChar[string(x)] + 1
		} else {
			countChar[string(x)] = 1
		}
	}
	for i,val := range countChar{
		fmt.Println("Count of char ", i, " is ", val, " in a name ",lower)
	}
	//fmt.Println(countChar)
}