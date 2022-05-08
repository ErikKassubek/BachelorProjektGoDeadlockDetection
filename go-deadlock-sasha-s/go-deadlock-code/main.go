package main

import (
	"time"

	"github.com/sasha-s/go-deadlock"
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
	var x deadlock.Mutex
	finished := make(chan bool)

	go func() {
		// first go routine
		x.Lock()
		time.Sleep(40 * time.Second)
		x.Unlock()
	}()

	go func() {
		// second go routine
		time.Sleep(2 * time.Second)
		x.Lock()
		x.Unlock()
		finished <- true
	}()

	<-finished
}

func test_4() {
	var x Mutex
	finished := make(chan bool)

	go func() {
		x.Lock()
		time.Sleep(40 * time.Second)
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
	var x deadlock.Mutex
	x.Lock()
	x.Lock()
	x.Unlock()
}
func test_6() {
	var x deadlock.Mutex
	var y deadlock.Mutex

	x.Lock()
	y.Lock()
	x.Lock()
	x.Unlock()
}

func main() {
	test_3()
}
