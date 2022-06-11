package main

import (
	deadlock "github.com/ErikKassubek/Deadlock-Go"
)

func main() {
	deadlock.SetCollectSingleLevelLockInformation(true)

	deadlock.Initialize()
	defer deadlock.FindPotentialDeadlocks()

	ch := make(chan bool, 4)

	// actualDeadlock(ch)
	// <-ch
	potentialDeadlock1(ch)
	<-ch
	potentialDeadlockThreeEdgeCirc(ch)
	<-ch
	noPotentialDeadlockGuardLocks(ch)
}
