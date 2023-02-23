//Go program to find single unique element in a slice
package main

import "fmt"

func main() {
	//y := []int{1, 2, 2, 3, 1}
	x := []int{2, 3, 4, 2, 4, 3, 5}

	for i := 0; i < len(x); i++ {
		for j := i + 1; j < len(x); j++ {
			if (x[i] ^ x[j]) == 0 {
				x[i] = 0
				x[j] = 0

			}

		}
		if x[i] != 0 {
			fmt.Println("Unique Element in a slice is", x[i])
		}
	}

	fmt.Println(x)

	mapApproach()
}

func mapApproach(){
	numbers:=[]int{2, 3, 4, 2, 4, 3, 5}
	unique:= make(map[int]int)
	for _, val:=range numbers{
		_,ok:=unique[val]
		if ok{
			unique[val] = unique[val] + 1
		}else{
			unique[val] = 1
		}
	}
	for val, count := range unique{
		if count == 1{
			fmt.Println("Unique number is: ", val)
		}
	}

}