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

	deadlock "github.com/ErikKassubek/Deadlock-Go"
	sasha "github.com/sasha-s/go-deadlock"
)

// simple example for potential deadlock
func SashaPotentialDeadlock(c chan<- bool) {
	var x sasha.Mutex
	var y sasha.Mutex
	var z sasha.Mutex
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
func SashaPotentialDeadlockThreeEdgeCirc(c chan<- bool) {
	var x sasha.Mutex
	var y sasha.Mutex
	var z sasha.Mutex

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
func SashaNoPotentialDeadlockGuardLocks(c chan<- bool) {
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

	c <- true

}

// actual deadlock
func SashaActualDeadlock(c chan<- bool) {
	var x sasha.Mutex
	var y sasha.Mutex
	var z sasha.Mutex
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

func SashaDoubleLogging(c chan<- bool) {
	var x sasha.Mutex
	ch := make(chan bool, 2)
	ch2 := make(chan bool)
	go func() {
		// deadlock.NewRoutine()
		x.Lock()
		x.Lock()
		ch2 <- true
		x.Unlock()
		ch <- true
	}()

	go func() {
		deadlock.NewRoutine()
		<-ch2
		ch <- true
	}()

	<-ch
	<-ch
	c <- true
}

func RunSasha(c chan<- bool) {
	ch := make(chan bool, 3)
	sasha.Opts.OnPotentialDeadlock = func() {}
	SashaPotentialDeadlock(ch)
	<-ch
	SashaPotentialDeadlockThreeEdgeCirc(ch)
	<-ch
	SashaNoPotentialDeadlockGuardLocks(ch)
	<-ch
	SashaDoubleLogging(ch)
	c <- true
}
