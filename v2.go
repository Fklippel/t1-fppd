// versao de: https://www.baeldung.com/cs/dining-philosophers adaptada para um mutex global
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// OUTPUT
//    N, THINKING, HUNGRY, EATING states, and LEFT, RIGHT index calculations.

const (
	N        = 5
	THINKING = 0
	HUNGRY   = 1
	EATING   = 2
)

func LEFT(i int) int { return (i + N - 1) % N }

func RIGHT(i int) int { return (i + 1) % N }

// OUTPUT
//    Initializes the synchronization primitives for the dining philosophers problem

var (
	state      [N]int
	semaphores [N]chan struct{}
	mutex      sync.Mutex
)

func SynchronizationPrimitives() {
	for i := 0; i < N; i++ {
		semaphores[i] = make(chan struct{}, 1)
	}
}

// INPUT
//    philosophernumber = the number representing the philosopher
// OUTPUT
//    the continuous routine of a philosopher alternating between thinking and eating

func PhilosopherRoutine(philosophernumber int) {
	i := philosophernumber

	for {
		Think(i)
		TakeForks(i)
		Eat(i)
		PutForks(i)
	}
}

// INPUT
//    i = index of the philosopher
// OUTPUT
//    updates the state of the philosopher and waits on the semaphore if conditions are met

func Check(i int) {
	if state[i] == HUNGRY && state[LEFT(i)] != EATING && state[RIGHT(i)] != EATING {
		state[i] = EATING
		semaphores[i] <- struct{}{}
	}
}

// INPUT
//    i = philosopher number
// OUTPUT
//    The philosopher i attempts to take the forks to start eating

func TakeForks(i int) {
//	mutex.Lock()
//	state[i] = HUNGRY
	fmt.Printf("Filosofo %d esta com FOME\n", i)
//	Check(i)
	mutex.Lock()
//	<-semaphores[i]
}

// INPUT
//    philosophernumber = the index of the philosopher
// OUTPUT
//    The philosopher puts down the forks and the system checks if neighbors can start eating

func PutForks(philosophernumber int) {
	i := philosophernumber
//	mutex.Lock()

//	state[i] = THINKING
	fmt.Printf("Filosofo %d largou os garfos\n", i)
//	Check(LEFT(i))
//	Check(RIGHT(i))

	mutex.Unlock()
}

// OUTPUT
//   Puts the philosopher thread to sleep for a random duration between 1 and 5

func Think(i int) {
	duration := rand.Intn(5) + 1
	fmt.Printf("Filosofo %d esta PENSANDO por %ds\n", i, duration)
	time.Sleep(time.Duration(duration) * time.Second)
}

// OUTPUT
//    Simulates eating by sleeping for a random duration between 1 and 3 seconds

func Eat(i int) {
	duration := rand.Intn(3) + 1
	fmt.Printf("Filosofo %d esta COMENDO por %ds\n", i, duration)
	time.Sleep(time.Duration(duration) * time.Second)
}

func main() {
	SynchronizationPrimitives()

	for i := 0; i < N; i++ {
		go PhilosopherRoutine(i)
	}

	select {}
}
