package main

import "fmt"

func main() {
	a := [...]int{0, 1, 2, 3}
	x := a[:1]
	fmt.Println(a)
	y := a[2:]
	fmt.Println(a)
	x = append(x, y...)
	fmt.Println(x, len(x), cap(x))
	fmt.Println(a, len(a),cap(a))
	x = append(x, y...)
	fmt.Println(a, x)
	//a := 0,1,2,3
	//x := 0,2,3,2,3

	println(bar())
}


func bar() (r int) {
	defer func() {
		r += 4
		if recover() != nil {
			r += 8
		}
	}()

	var f func()
	defer f()
	f = func() {
		r += 2
	}

	return 1
}
/*
package main

import (
	"fmt"
	"sync"
)
var result = make(chan []string)

func FuncOne(nc chan []string, wg *sync.WaitGroup){
	defer wg.Done()
	//fmt.Println(<-nc)
	y := <-nc
	result <- y[:(len(y)-1)]
	close(nc)
}

func main(){
	name := "PUNEET"
	wg := sync.WaitGroup{}
	arr:=[]string{}
	nameChan := make(chan []string)
	wg.Add(1)
	go FuncOne(nameChan,&wg)
	for _, val := range name{
		arr = append(arr, string(val))
	}
	nameChan<-arr
	fmt.Println(<-result)
	wg.Wait()
}

func main() {
	x := "PUNEET"
	var arr []string
	for _, val := range x {
		arr = append(arr, string(val))
	}
	lenght := len(arr)
	fmt.Println(arr[:(lenght - 1)])
}
*/
/*
package main

import "fmt"

// recover function to handle panic
func handlePanic() {

	// detect if panic occurs or not
	a := recover()

	if a != nil {
		fmt.Println("RECOVER", a)
	}

}

func division(num1, num2 int) {

	// execute the handlePanic even after panic occurs
	defer handlePanic()

	// if num2 is 0, program is terminated due to panic
	if num2 == 0 {
		panic("Cannot divide a number by zero")
		fmt.Println("panic Result: ", num1, num2)
	} else {
		result := num1 / num2
		fmt.Println("Result: ", num1, num2, result)
	}

}

func main() {

	division(4, 2) //2
	division(8, 0) //
	division(2, 8) //0.25

}

package main

import "fmt"

func main() {
	var quit chan bool
	go func() {
		for {
			select {
			case <-quit:
				fmt.Println("quit")
				return
			default:
				fmt.Println("default")
			}
		}
	}()
	quit <- true
}



package main

import (
	"fmt"
)

func main() {
	c := make(chan string)
	go countCat(c)

	for i := 0; i < 6; i++ {
		message := <-c
		fmt.Println(message)
	}
}

func countCat(c chan string) {
	for i := 0; i < 5; i++ {
		c <- "Cat"
	}
	close(c)
}


package main

import "fmt"

func main() {
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}

 */

/*
package main

import "fmt"

const N = 4

func main(){
	m := make(map[int]*int)
	for i:=1;i<=N;i++{
		m[i]=&i
		fmt.Println(&i)
	}

	for key, value := range m{
		fmt.Println(key,*value)
	}

	//fmt.Println(m)
}
*/

/*
package main

import (
	"fmt"
)

func main() {
	s1 := []int{10, 20, 30, 40, 50}
	s2 := s1[2:4]        //{30,40}
	fmt.Println(len(s2)) //2
	fmt.Println(cap(s2)) //3
	s2 = append(s2, 60)  //{30,40,60}
	fmt.Println(len(s2)) //3
	fmt.Println(cap(s2)) //3
	//s2 = append(s2, 70)  //{30,40,60,70} len-4 cap-6
	fmt.Println(s1) //{}
	fmt.Println(s2) //{30,40,60}
}
*/

/*

*/
