package main

import (
	"fmt"
	"sync"
	"time"
)

func test_1() {

	Opts.PrintAllCurrentGoroutines = false
	var x Mutex
	var y Mutex

	go func() {
		y.Lock()
		x.Lock()
		x.Unlock()
		y.Unlock()
	}()
	go func() {
		x.Lock()
		y.Lock()
		y.Unlock()
		x.Unlock()
	}()

	time.Sleep(time.Second * 3)
}

func test_2() {
	var x Mutex
	var y Mutex

	go func() {
		y.Lock()
		x.Lock()
		x.Unlock()
		y.Unlock()
		x.Lock()
		y.Lock()
		y.Unlock()
		y.Unlock()
	}()
}

// no possible deadlock but still error message
func test_3() {
	Opts.DeadlockTimeout = 10 * time.Second
	Opts.OnPotentialDeadlock = func() {
		fmt.Println("func")
	}

	var x Mutex
	finished := make(chan bool)

	go func() {
		x.Lock()
		time.Sleep(20 * time.Second)
		x.Unlock()
		finished <- true
	}()

	go func() {
		time.Sleep(1 * time.Second)
		x.Lock()
		x.Unlock()
	}()

	<-finished
}

func test_4() {
	Opts.DeadlockTimeout = 10 * time.Second
	var x Mutex
	finished := make(chan bool)

	go func() {
		x.Lock()
		time.Sleep(4 * time.Second)
		x.Unlock()
		finished <- true
	}()

	go func() {
		time.Sleep(1 * time.Second)
		x.Lock()
		x.Unlock()
	}()

	time.Sleep(7 * time.Second)
	x.Lock()
	time.Sleep(10 * time.Second)
	x.Unlock()

	<-finished
}

func test_5() {
	var mu sync.Mutex
	mu.Lock()
	mu.Lock()
	mu.Unlock()
}

func main() {
	test_5()
}
