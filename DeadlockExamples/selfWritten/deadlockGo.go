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
	"math/rand"
	"time"

	deadlock "github.com/ErikKassubek/Deadlock-Go"
)

// ========== Mutex ============

// ------Lock-------

// 1. simple example for potential deadlock with two routines
func DeadlockGoPotentialDeadlock() {
	x := deadlock.NewLock()
	y := deadlock.NewLock()
	z := deadlock.NewLock()
	ch := make(chan bool, 2)

	go func() {
		z.Lock()
		z.Unlock()
		time.Sleep(time.Second)
		x.Lock()
		y.Lock()
		y.Unlock()
		x.Unlock()

		ch <- true
	}()

	go func() {
		y.Lock()
		x.Lock()
		x.Unlock()
		y.Unlock()

		ch <- true
	}()

	<-ch
	<-ch
}

// 2. test with 3 edge loop
func DeadlockGoPotentialDeadlockThreeEdgeCirc() {
	x := deadlock.NewLock()
	y := deadlock.NewLock()
	z := deadlock.NewLock()

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

// 3. no deadlock because of Gate locks
func DeadlockGoNoPotentialDeadlockGateLocks() {
	x := deadlock.NewLock()
	y := deadlock.NewLock()
	z := deadlock.NewLock()
	ch := make(chan bool, 2)

	go func() {
		z.Lock()
		x.Lock()
		y.Lock()
		time.Sleep(time.Second * time.Duration(rand.Float64()))
		y.Unlock()
		x.Unlock()
		z.Unlock()

		ch <- true
	}()

	go func() {
		z.Lock()
		y.Lock()
		x.Lock()
		time.Sleep(time.Second * time.Duration(rand.Float64()))
		x.Unlock()
		y.Unlock()
		z.Unlock()

		ch <- true
	}()

	<-ch
	<-ch
}

// 4. deadlock with nested go routines
func DeadlockGoNestedRoutines() {
	x := deadlock.NewLock()
	y := deadlock.NewLock()
	ch := make(chan bool)
	ch2 := make(chan bool)

	go func() {
		x.Lock()
		go func() {
			y.Lock()
			y.Unlock()
			ch <- true
		}()
		<-ch
		x.Unlock()
		ch2 <- true
	}()
	go func() {
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
func DeadlockGoDoubleLogging() {
	x := deadlock.NewLock()
	ch := make(chan bool, 2)
	ch2 := make(chan bool)
	go func() {
		x.Lock()
		x.Lock()
		ch2 <- true
		x.Unlock()
		ch <- true
	}()

	go func() {
		<-ch2
		ch <- true
	}()

	<-ch
	<-ch
}

// 6.1 actual deadlock with two routines
func DeadlockGoActualDeadlock() {
	x := deadlock.NewLock()
	y := deadlock.NewLock()
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
func DeadlockGoActualDeadlockThree() {
	x := deadlock.NewLock()
	y := deadlock.NewLock()
	z := deadlock.NewLock()

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

// -------------- trylock --------------

// 7. double locking including trylock
func DoubleLockingIncludingTryLock() {
	x := deadlock.NewLock()
	x.TryLock()
	x.Lock()
}

// 8. deadlock including tryLock
func DeadlockIncludingTryLock() {
	x := deadlock.NewLock()
	y := deadlock.NewLock()

	ch := make(chan bool, 2)

	go func() {
		time.Sleep(time.Second) // remove for actual deadlock
		a := x.TryLock()
		y.Lock()
		y.Unlock()
		if a {
			x.Unlock()
		}
		ch <- true
	}()

	go func() {
		a := y.TryLock()
		x.Lock()
		if a {
			y.Unlock()
		}
		x.Unlock()
		ch <- true
	}()

	<-ch
	<-ch
}

// =========== RW-Lock ==========

// 9. can be used to test all combinations or RLock and  Lock
func DeadlockRwDeadlock() {
	x := deadlock.NewRWLock()
	y := deadlock.NewRWLock()

	ch := make(chan bool, 2)

	go func() {
		x.RLock()
		y.RLock()
		y.RUnlock()
		x.RUnlock()

		ch <- true
	}()

	go func() {
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
func DeadlockGateLocksRW() {
	x := deadlock.NewRWLock()
	y := deadlock.NewRWLock()
	z := deadlock.NewRWLock()
	ch := make(chan bool, 2)

	go func() {
		z.RLock()
		x.Lock()
		y.Lock()
		y.Unlock()
		x.Unlock()
		z.RUnlock()

		ch <- true
	}()

	go func() {
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

// 11. Double Locking with rwLocks
func DeadlockRWDoubleLogging() {
	x := deadlock.NewRWLock()
	x.RLock()
	x.RLock()
	x.RUnlock()
	x.RUnlock()
}

func RunDeadlockGo() {
	// DeadlockGoPotentialDeadlock()
	// DeadlockGoPotentialDeadlockThreeEdgeCirc()
	// DeadlockGoNoPotentialDeadlockGateLocks()
	// DeadlockGoNestedRoutines()
	// DeadlockGoDoubleLogging()
	// DeadlockGoActualDeadlock()
	// DeadlockGoActualDeadlockThree()
	DoubleLockingIncludingTryLock()
	// DeadlockIncludingTryLock()
	// DeadlockRwDeadlock()
	// DeadlockGateLocksRW()
	// DeadlockRWDoubleLogging()
}
