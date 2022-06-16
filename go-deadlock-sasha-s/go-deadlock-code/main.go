package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/sasha-s/go-deadlock"
)

func test_1() {
	var x deadlock.Mutex
	var y deadlock.Mutex
	ch := make(chan bool, 2)

	go func() {
		y.Lock()
		x.Lock()
		x.Unlock()
		y.Unlock()
		ch <- true
	}()
	go func() {
		x.Lock()
		y.Lock()
		y.Unlock()
		x.Unlock()
		ch <- true
	}()

	<-ch
	<-ch
}

func test_2() {
	var x deadlock.Mutex
	var y deadlock.Mutex

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
	var x deadlock.Mutex
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

// does not detect
func test_7() {
	var x deadlock.Mutex
	var y deadlock.Mutex
	ch := make(chan bool)

	go func() {
		x.Lock()
		go func() {
			y.Lock()
			y.Unlock()
			ch <- true
		}()
		<-ch
		x.Unlock()
	}()

	go func() {
		y.Lock()
		x.Lock()
		x.Lock()
		y.Lock()
	}()
}

func test_8() {
	var x deadlock.Mutex
	var y deadlock.Mutex
	var z deadlock.Mutex
	ch := make(chan bool)

	go func() {
		time.Sleep(time.Second)
		x.Lock()
		y.Lock()
		z.Lock()
		z.Unlock()
		y.Unlock()
		x.Unlock()
		ch <- true
	}()

	z.Lock()
	y.Lock()
	x.Lock()
	x.Unlock()
	y.Unlock()
	z.Unlock()

	<-ch

}

func test_9() {
	var x deadlock.Mutex
	var y deadlock.Mutex

	x.Lock()
	y.Lock()
	y.Unlock()
	x.Unlock()

	y.Lock()
	x.Lock()
	x.Unlock()
	y.Unlock()
}

func test_10() {
	var x deadlock.Mutex
	var y deadlock.Mutex
	var z deadlock.Mutex
	ch := make(chan bool, 3)

	go func() {
		x.Lock()
		y.Lock()
		y.Unlock()
		x.Unlock()
		ch <- true
	}()

	go func() {
		y.Lock()
		z.Lock()
		z.Unlock()
		y.Unlock()
		ch <- true
	}()

	go func() {
		z.Lock()
		x.Lock()
		x.Unlock()
		z.Unlock()
		ch <- true
	}()

	<-ch
	<-ch
	<-ch
}

func test_11() {
	var x sync.Mutex
	ch := make(chan bool)

	go func() {
		x.Lock()
		x.Lock()
		fmt.Println()
		x.Unlock()
		ch <- true
	}()

	<-ch
}

// guard locks
func test_12() {
	var x deadlock.Mutex
	var y deadlock.Mutex
	var z deadlock.Mutex

	ch := make(chan bool, 2)

	go func() {
		z.Lock()
		y.Lock()
		x.Lock()
		x.Unlock()
		y.Unlock()
		z.Unlock()
		ch <- true
	}()
	go func() {
		z.Lock()
		x.Lock()
		y.Lock()
		y.Unlock()
		x.Unlock()
		z.Unlock()
		ch <- true
	}()

	<-ch
	<-ch
}

func main() {
	fmt.Println("run")
	test_7()
	time.Sleep(time.Second)
}
