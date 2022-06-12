package main

import (
	"fmt"
	"time"

	deadlock "github.com/ErikKassubek/Deadlock-Go"
	sasha "github.com/sasha-s/go-deadlock"
)

func own(c chan<- bool) {
	deadlock.SetCollectSingleLevelLockInformation(true)
	// deadlock.SetCollectCallStack(true)

	deadlock.Initialize()
	defer deadlock.FindPotentialDeadlocks()

	ch := make(chan bool, 200)

	deadlockGoPotentialDeadlock1(ch)
	<-ch
	deadlockGoPotentialDeadlockThreeEdgeCirc(ch)
	<-ch
	deadlockGoNoPotentialDeadlockGuardLocks(ch)
	<-ch
	// deadlockGoActualDeadlock(ch)
	// <-ch

	close(ch)
	c <- true
}

func sacha(c chan<- bool) {
	sasha.Opts.OnPotentialDeadlock = func() {}
	ch := make(chan bool, 200)

	sashaPotentialDeadlock1(ch)
	<-ch
	sashaPotentialDeadlockThreeEdgeCirc(ch)
	<-ch
	sashaNoPotentialDeadlockGuardLocks(ch)
	<-ch
	// sashaActualDeadlock(ch)
	// <-ch

	close(ch)
	c <- true
}

func main() {
	runOwn := true
	runSahsa := true
	showTime := true
	c := make(chan bool, 2)

	var durationOwn time.Duration
	var durationSasha time.Duration

	if runOwn {
		startOwn := time.Now()
		own(c)
		durationOwn = time.Since(startOwn)
		<-c
	}

	if runSahsa {
		startSasha := time.Now()
		sacha(c)
		durationSasha = time.Since(startSasha)
		<-c
	}

	if showTime {
		if runOwn {
			fmt.Println("Time for Deadlock-Go: ", durationOwn)
		}
		if runSahsa {
			fmt.Println("Time for Sasha:       ", durationSasha)
		}
	}
}
