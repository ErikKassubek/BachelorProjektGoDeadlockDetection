package selfWritten

/*
Author: Erik Kassubek <erik-kassubek@t-online.de>
Date: 2022-06-12
*/

/*
examplesSahsa.go
This file implements examples for testing deadlocks with the deadlock-go
(https://github.com/ErikKassubek/Deadlock-Go) tool. These examples are the same
as in the examplesSasha.go file
*/

import (
	"sync"
)

// ========== Mutex ============

// ------Lock-------

// 1. simple example for potential deadlock with two routines
func PotentialDeadlock() {
	var x sync.Mutex
	var y sync.Mutex
	ch := make(chan bool, 2)
	ch2 := make(chan bool)

	go func() {
		<-ch2
		x.Lock()
		y.Lock()
		y.Unlock()
		x.Unlock()
		ch <- true
	}()

	go func() {
		y.Lock()
		x.Lock()
		ch2 <- true
		x.Unlock()
		y.Unlock()
		ch <- true
	}()

	<-ch
	<-ch
}

// 2. test with 3 edge loop
func PotentialDeadlockThreeEdgeCirc() {
	var x sync.Mutex
	var y sync.Mutex
	var z sync.Mutex

	ch := make(chan bool, 3)
	ch2 := make(chan bool)
	ch3 := make(chan bool)

	go func() {
		x.Lock()
		y.Lock()
		ch2 <- true
		y.Unlock()
		x.Unlock()

		ch <- true
	}()

	go func() {
		<-ch2
		y.Lock()
		z.Lock()
		ch3 <- true
		z.Unlock()
		y.Unlock()

		ch <- true
	}()

	go func() {
		<-ch3
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

// 3. no deadlock because of Gate locks
func NoPotentialDeadlockGateLocks() {
	var x sync.Mutex
	var y sync.Mutex
	var z sync.Mutex
	ch := make(chan bool, 2)

	go func() {
		z.Lock()
		x.Lock()
		y.Lock()
		y.Unlock()
		x.Unlock()
		z.Unlock()

		ch <- true
	}()

	go func() {
		z.Lock()
		y.Lock()
		x.Lock()
		x.Unlock()
		y.Unlock()
		z.Unlock()

		ch <- true
	}()

	<-ch
	<-ch
}

// 4. deadlock with nested go routines
func NestedRoutines() {
	var x sync.Mutex
	var y sync.Mutex
	ch := make(chan bool)
	ch2 := make(chan bool)
	ch3 := make(chan bool) // prevent actual deadlock

	go func() {
		x.Lock()
		go func() {
			y.Lock()
			y.Unlock()
			ch <- true
		}()
		<-ch
		x.Unlock()
		ch3 <- true
		ch2 <- true
	}()
	go func() {
		<-ch3
		y.Lock()
		x.Lock()
		x.Unlock()
		y.Unlock()
		ch2 <- true
	}()

	<-ch2
	<-ch2

}

// =========== RW-Lock ==========

// 9. can be used to test all combinations or RLock and  Lock
func RwDeadlock() {
	var x sync.RWMutex
	var y sync.RWMutex

	ch := make(chan bool, 2)
	ch2 := make(chan bool) // prevent actual deadlock

	go func() {
		x.RLock()
		y.RLock()
		ch2 <- true
		y.RUnlock()
		x.RUnlock()

		ch <- true
	}()

	go func() {
		<-ch2
		y.RLock()
		x.RLock()
		x.RUnlock()
		y.RUnlock()

		ch <- true
	}()

	<-ch
	<-ch
}

// 10. Can be used to check RW-Locks as gate locks (both Lock and R-Lock)
func GateLocksRW() {
	var x sync.RWMutex
	var y sync.RWMutex
	var z sync.RWMutex
	ch := make(chan bool, 2)
	ch2 := make(chan bool) // prevent actual deadlock

	go func() {
		z.RLock()
		x.Lock()
		y.Lock()
		ch2 <- true
		y.Unlock()
		x.Unlock()
		z.RUnlock()

		ch <- true
	}()

	go func() {
		<-ch2
		z.RLock()
		y.Lock()
		x.Lock()
		x.Unlock()
		y.Unlock()
		z.RUnlock()

		ch <- true
	}()

	<-ch
	<-ch
}

func RunWithoutDetector() {
	for i := 0; i < 1000; i++ {
		PotentialDeadlock()
		// PotentialDeadlockThreeEdgeCirc()
		// NoPotentialDeadlockGateLocks()
		// NestedRoutines()
		// RwDeadlock()
		// GateLocksRW()

	}
}
