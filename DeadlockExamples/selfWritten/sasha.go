package selfWritten

/*
Author: Erik Kassubek <erik-kassubek@t-online.de>
Date: 2022-06-12
*/

/*
examplesSahsa.go
This file implements examples for testing deadlocks with the sasha-s/go-deadlock
(https://github.com/sasha-s/go-deadlock) tool. These examples are the same as
in the examplesDeadlockGo.go file
*/

import (
	"time"

	sasha "github.com/sasha-s/go-deadlock"
)

// =========== Mutex ===========

// --------- Lock ---------

// 1. simple example for potential deadlock
func SashaPotentialDeadlock() {
	var x sasha.Mutex
	var y sasha.Mutex
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
func SashaPotentialDeadlockThreeEdgeCirc() {
	var x sasha.Mutex
	var y sasha.Mutex
	var z sasha.Mutex

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

// 3. no deadlock because of gate locks
func SashaNoPotentialDeadlockGateLocks() {
	var x sasha.Mutex
	var y sasha.Mutex
	var z sasha.Mutex
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

// 4. nested routines
func SashaNestedRoutines() {
	var x sasha.Mutex
	var y sasha.Mutex
	ch := make(chan bool)
	ch2 := make(chan bool, 2)
	ch3 := make(chan bool) // prevents actual deadlock

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

// 5. double locking
func SashaDoubleLogging() {
	var x sasha.Mutex
	ch := make(chan bool, 2)
	ch2 := make(chan bool)
	go func() {
		x.Lock()
		x.Lock()
		ch2 <- true
		x.Unlock()
		ch <- true
	}()

	<-ch
	<-ch
}

// 6.1 actual deadlock with two routines
func SashaActualDeadlock() {
	var x sasha.Mutex
	var y sasha.Mutex
	ch := make(chan bool, 2)
	ch2 := make(chan bool)

	go func() {
		x.Lock()
		time.Sleep(time.Second)
		ch2 <- true
		y.Lock()
		y.Unlock()
		x.Unlock()
		ch <- true
	}()

	go func() {
		y.Lock()
		<-ch2
		x.Lock()
		x.Unlock()
		y.Unlock()
		ch <- true
	}()

	<-ch
	<-ch
}

// 6.2 actual deadlock with tree routines
func SashaGoActualDeadlockThree() {
	var x sasha.Mutex
	var y sasha.Mutex
	var z sasha.Mutex

	ch := make(chan bool, 3)
	ch2 := make(chan bool)
	ch3 := make(chan bool)
	ch4 := make(chan bool)
	ch5 := make(chan bool)

	go func() {
		x.Lock()
		ch2 <- true
		<-ch4
		y.Lock()
		y.Unlock()
		x.Unlock()
		ch <- true
	}()

	go func() {
		<-ch2
		y.Lock()
		ch3 <- true
		ch4 <- true
		<-ch5
		z.Lock()
		x.Unlock()
		y.Unlock()
		ch <- true
	}()

	go func() {
		<-ch3
		z.Lock()
		ch5 <- true
		x.Lock()
		x.Unlock()
		y.Unlock()
		ch <- true
	}()

	<-ch
	<-ch
	<-ch
}

// go-deadlock has no try-locks. Therefor there are no implementations
// for 7, 8

// ======== RW-Locks ========
// 9 can be used to test all combinations or RLock and  Lock
func SashaRwDeadlock() {
	var x sasha.RWMutex
	var y sasha.RWMutex

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
func SashaGateLocksRW() {
	var x sasha.RWMutex
	var y sasha.RWMutex
	var z sasha.RWMutex
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

// 11. Double Locking with rw-Locks
func SashaRWDoubleLogging() {
	var x sasha.RWMutex
	x.RLock()
	x.RLock()
	x.RUnlock()
	x.RUnlock()
}

func RunSasha() {
	sasha.Opts.OnPotentialDeadlock = func() {} // for timing analysis
	for i := 0; i < 1000; i++ {
		// SashaPotentialDeadlock()
		// SashaPotentialDeadlockThreeEdgeCirc()
		// SashaNoPotentialDeadlockGateLocks()
		// SashaNestedRoutines()
		// SashaDoubleLogging()
		// SashaActualDeadlock()
		// SashaGoActualDeadlockThree()
		SashaRwDeadlock()
		// SashaGateLocksRW()
		// SashaRWDoubleLogging()
	}
}
