package main

import "fmt"

/*func TwoSumProblem(x []int, target int)(int, int){
	sum := 0
	for i:=0;i<len(x);i++{
		for j:= i+1;j<len(x);j++{
			sum = x[i]+x[j]
			if sum == target{
				return i,j
				break
			}
		}

	}
	return -1,-1
}*/

func TwoSumProblem(nums []int, target int)(interface{},interface{}){
	seenNums := map[int]int{}
	for i, num := range nums {
		potentialMatch := target - num
		j, found := seenNums[potentialMatch]
		fmt.Println(i, j, seenNums[potentialMatch])
		if found {
			return i, j
		}
		seenNums[num] = i
		fmt.Println(seenNums)
	}
	return nil,nil
}

func main(){
	x := []int{3,2,4}
	target := 6
	i, j := TwoSumProblem(x, target)
	fmt.Println(i,j)
}