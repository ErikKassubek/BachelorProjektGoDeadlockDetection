package main

import (
	"math/rand"
	"time"

	deadlock "github.com/ErikKassubek/Deadlock-Go"
)

func potentialDeadlock1() {
	deadlock.Initialize()
	defer deadlock.Finalize()

	x := deadlock.NewLock()
	y := deadlock.NewLock()
	z := deadlock.NewLock()
	ch := make(chan bool, 2)

	go func() {
		deadlock.NewRoutine()

		z.Lock()
		z.Unlock()
		x.Lock()
		y.Lock()
		time.Sleep(time.Second * time.Duration(rand.Float64()))
		y.Unlock()
		x.Unlock()

		ch <- true
	}()

	go func() {
		deadlock.NewRoutine()

		y.Lock()
		x.Lock()
		time.Sleep(time.Second * time.Duration(rand.Float64()))
		x.Unlock()
		y.Unlock()

		ch <- true
	}()

	<-ch
	<-ch

}

// test with 3 edge loop
func potentialDeadlockThreeEdgeCirc() {
	deadlock.Initialize()
	defer deadlock.Finalize()

	x := deadlock.NewLock()
	y := deadlock.NewLock()
	z := deadlock.NewLock()

	ch := make(chan bool, 3)

	go func() {
		deadlock.NewRoutine()

		x.Lock()
		y.Lock()
		time.Sleep(time.Second * time.Duration(rand.Float64()))
		y.Unlock()
		x.Unlock()

		ch <- true
	}()

	go func() {
		deadlock.NewRoutine()

		y.Lock()
		z.Lock()
		time.Sleep(time.Second * time.Duration(rand.Float64()))
		z.Unlock()
		y.Unlock()

		ch <- true
	}()

	go func() {
		deadlock.NewRoutine()
		z.Lock()
		x.Lock()
		time.Sleep(time.Second * time.Duration(rand.Float64()))
		x.Unlock()
		z.Unlock()

		ch <- true
	}()

	<-ch
	<-ch
	<-ch

}

func potentialDeadlockGuardLocks() {
	deadlock.Initialize()
	defer deadlock.Finalize()

	x := deadlock.NewLock()
	y := deadlock.NewLock()
	z := deadlock.NewLock()
	ch := make(chan bool, 2)

	go func() {
		deadlock.NewRoutine()

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
		deadlock.NewRoutine()

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

func actualDeadlock() {
	deadlock.Initialize()

	x := deadlock.NewLock()
	y := deadlock.NewLock()
	z := deadlock.NewLock()
	ch := make(chan bool, 2)
	ch2 := make(chan bool)

	go func() {
		deadlock.NewRoutine()
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
		deadlock.NewRoutine()
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

func main() {
	potentialDeadlock1()
}
