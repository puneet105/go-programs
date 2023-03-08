package main

import (
	"fmt"
	"sync"
)

func PrintEven(wg *sync.WaitGroup, done chan bool) {
	defer wg.Done()
	for i:=0;i<20;i++{
		if i%2 == 0{
			fmt.Printf("EVEN:%d ",i)

		}
		done<-true
	}
	close(done)
}

func PrintOdd(wg *sync.WaitGroup,  done chan bool) {
	defer wg.Done()
	for i:=0;i<20;i++{
		<-done
		if i%2 != 0{
			fmt.Printf("ODD:%d ",i)
		}
	}

}

func main() {
	var wg sync.WaitGroup
	done := make(chan bool)
	wg.Add(2)
	go PrintEven(&wg,done)
	go PrintOdd( &wg,done)
	wg.Wait()
}

/*package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var done int32
	syncChannel := make(chan bool) // unbuffered channel.

	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {
		// prints even numbers.
		defer wg.Done()

		for i := 0; i < 50; i += 2 {
			<-syncChannel

			fmt.Printf("Even: %d\n", i)

			syncChannel <- true
		}
		atomic.StoreInt32(&done, 1)
	}()

	syncChannel <- true

	go func() {
		// prints odd numbers.
		defer wg.Done()

		for i := 1; i < 50; i += 2 {
			<-syncChannel

			fmt.Printf("Odd:%d\n", i)

			if atomic.LoadInt32(&done) != 0 {
				return
			}

			syncChannel <- true
		}
	}()

	wg.Wait()
	close(syncChannel)
}*/
