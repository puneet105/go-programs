// Q2: Solve the "Dining Philosophers Problem" using Goroutines and channels.
// How It Works
// 1. Forks as Channels: Each fork is a buffered channel with capacity 1, representing its availability.
// 2. Host Channel: Ensures at most N-1 philosophers eat at the same time, avoiding deadlocks.
// 3. Lifecycle:
//     * Each philosopher thinks.
//     * Requests permission from host.
//     * Picks up forks (left then right).
//     * Eats for a random time.
//     * Releases forks & notifies host when done.
//     * Repeats for 3 cycles before exiting.


package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Philosopher struct
type Philosopher struct {
	id              int
	leftFork, rightFork chan struct{}
}

// Think simulates thinking
func (p Philosopher) Think() {
	fmt.Printf("Philosopher %d is thinking...\n", p.id)
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
}

// Eat simulates eating
func (p Philosopher) Eat(host chan struct{}) {
	<-host // Request permission to eat
	fmt.Printf("Philosopher %d is hungry and trying to pick up forks...\n", p.id)

	<-p.leftFork
	fmt.Printf("Philosopher %d picked up left fork\n", p.id)

	<-p.rightFork
	fmt.Printf("Philosopher %d picked up right fork and starts eating\n", p.id)

	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))

	// Release forks
	p.rightFork <- struct{}{}
	p.leftFork <- struct{}{}
	fmt.Printf("Philosopher %d finished eating and released forks\n", p.id)

	host <- struct{}{} // Notify host that eating is done
}

// Start philosopher's lifecycle
func (p Philosopher) Start(wg *sync.WaitGroup, host chan struct{}) {
	defer wg.Done()
	for i := 0; i < 3; i++ { // Each philosopher eats 3 times
		p.Think()
		p.Eat(host)
	}
	fmt.Printf("Philosopher %d is done.\n", p.id)
}

func main() {
	const numPhilosophers = 5

	// Forks are represented as buffered channels (1 token each)
	forks := make([]chan struct{}, numPhilosophers)
	for i := range forks {
		forks[i] = make(chan struct{}, 1)
		forks[i] <- struct{}{} // Each fork is initially available
	}

	// Host channel to limit eating philosophers (N-1 rule)
	host := make(chan struct{}, numPhilosophers-1)
	for i := 0; i < numPhilosophers-1; i++ {
		host <- struct{}{}
	}

	// Initialize philosophers
	var wg sync.WaitGroup
	philosophers := make([]Philosopher, numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		philosophers[i] = Philosopher{
			id:        i,
			leftFork:  forks[i],
			rightFork: forks[(i+1)%numPhilosophers],
		}
	}

	// Start philosophers as goroutines
	for _, p := range philosophers {
		wg.Add(1)
		go p.Start(&wg, host)
	}

	wg.Wait()
	fmt.Println("All philosophers have finished dining.")
}