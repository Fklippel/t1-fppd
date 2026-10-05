// solucao de: https://www.youtube.com/watch?v=dCwhX-i8qT8
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numPhilosophers = 5
	maxThinkingTime = 3 * time.Second
	maxEatingTime   = 1 * time.Second
)

type philosopher struct {
	id                  int
	leftFork, rightFork chan struct{}
}

func (p *philosopher) think() {
	fmt.Printf("Philosopher %d is thinking\n", p.id)
	time.Sleep(time.Duration(rand.Intn(int(maxThinkingTime / time.Millisecond))) * time.Millisecond)
}

func (p *philosopher) eat(wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		if pickUpForks(p.leftFork, p.rightFork) {
			fmt.Printf("Philosopher %d is eating\n", p.id)
			time.Sleep(time.Duration(rand.Intn(int(maxEatingTime / time.Millisecond))) * time.Millisecond)
			putDownForks(p.leftFork, p.rightFork)
			p.think()
			return
		}
	}
}

func pickUpForks(leftFork, rightFork chan struct{}) bool {
	select {
	case leftFork <- struct{}{}:
		select {
		case rightFork <- struct{}{}:
			return true
		default:
			<-leftFork
			return false
		}
	default:
		return false
	}
}

func putDownForks(leftFork, rightFork chan struct{}) {
	<-leftFork
	<-rightFork
}

func main() {
	rand.Seed(time.Now().UnixNano())

	forks := make([]chan struct{}, numPhilosophers)
	for i := range forks {
		forks[i] = make(chan struct{}, 1)
	}

	philosophers := make([]*philosopher, numPhilosophers)
	for i := range philosophers {
		leftFork := forks[i]
		rightFork := forks[(i+1)%numPhilosophers]
		philosophers[i] = &philosopher{id: i, leftFork: leftFork, rightFork: rightFork}
	}

	var wg sync.WaitGroup
	wg.Add(numPhilosophers)

	for _, p := range philosophers {
		go p.eat(&wg)
	}

	wg.Wait()
	fmt.Println("All philosophers have finished eating")
}
