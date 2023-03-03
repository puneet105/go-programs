/*Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].*/

package main

import "fmt"

func main(){
	nums := []int{3,15,0,9}
	target := 9
	/*for i:=0;i<len(nums);i++{
		for j:=i+1;j<len(nums);j++{
				if (nums[i] + nums[j]) == target{
					fmt.Println(i,j)
					break
				}
		}
	}*/

	numMap := make(map[int]int)
	for i, num := range nums{
		x := target - num
		//fmt.Println(x)
		if j, match := numMap[x]; match{
			fmt.Println(i , j)
		}
		numMap[num] = i
		//fmt.Println(numMap)
	}


}
