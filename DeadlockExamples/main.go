package main

import (
	"DeadlockExamples/selfWritten"

	deadlock "github.com/ErikKassubek/Deadlock-Go"
)

func main() {
	c := make(chan bool, 2)
	deadlock.SetDoubleLockingDetection(true)
	deadlock.Initialize()

	// selfWritten.RunDeadlockGo(c)
	// <-c

	selfWritten.RunSasha(c)
	<-c
}
