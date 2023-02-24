// Array = [10,20,30,40,50,60,70,80], right rotate an array by k=3 times.
// k=1 (no. of time you want to rotate)

package main 

import "fmt"

func main(){
  intArr := []int{10,20,30,40,50,60,70,80}
  k:=1
  //rotate right
  for i :=0 ; i<k ; i++{
    lastElem := intArr[len(intArr) -1]
    for j := len(intArr) -1 ; j>0 ; j--{
      intArr[j] = intArr[j-1]
    }  
    intArr[0]=lastElem
  }
  fmt.Println(intArr)
  
  //rotate left
  intArr := []int{10,20,30,40,50,60,70,80}
	k:=1
	for i :=0 ; i<k ; i++{
		j := 0
		firstElem := intArr[0]
		for ; j<len(intArr)-1 ; j++{
			intArr[j] = intArr[j+1]
		}
		intArr[j]=firstElem
	}
	fmt.Println(intArr)
  
}
