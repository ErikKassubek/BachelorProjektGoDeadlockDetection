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

// simple example for potential deadlock
func DeadlockGoPotentialDeadlock(c chan<- bool) {
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

	c <- true

}

// test with 3 edge loop
func DeadlockGoPotentialDeadlockThreeEdgeCirc(c chan<- bool) {
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

	c <- true

}

// no deadlock because of guard locks
func DeadlockGoNoPotentialDeadlockGuardLocks(c chan<- bool) {
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

	c <- true

}

func DeadlockGoNestedRoutines(c chan<- bool) {
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

	c <- true
}

// actual deadlock
func DeadlockGoActualDeadlock(c chan<- bool) {
	x := deadlock.NewLock()
	y := deadlock.NewLock()
	z := deadlock.NewLock()
	ch := make(chan bool, 2)
	ch2 := make(chan bool)

	go func() {
		z.Lock()
		z.Unlock()
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

	c <- true
}

func DeadlockGoDoubleLogging(c chan<- bool) {
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
	c <- true
}

func RunDeadlockGo() {
	ch := make(chan bool, 6)
	// DeadlockGoPotentialDeadlock(ch)
	// <-ch
	// DeadlockGoPotentialDeadlockThreeEdgeCirc(ch)
	// <-ch
	// DeadlockGoNoPotentialDeadlockGuardLocks(ch)
	// <-ch
	DeadlockGoNestedRoutines(ch)
	<-ch
	// DeadlockGoDoubleLogging(ch)
	// <-ch
	// DeadlockGoActualDeadlock(ch)
	//<-ch
}
