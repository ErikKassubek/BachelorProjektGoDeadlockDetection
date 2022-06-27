package main

import (
	"DeadlockExamples/selfWritten"
	"time"

	deadlock "github.com/ErikKassubek/Deadlock-Go"
)

func main() {
	// deadlock.SetDoubleLockingDetection(true)

	defer deadlock.FindPotentialDeadlocks()

	selfWritten.RunDeadlockGo()

	// selfWritten.RunSasha()

	// otherExamples.RunCockroach3710()
	// otherExamples.RunCockroach9935()
	// otherExamples.RunCockroach6181()
	// otherExamples.RunCockroach7504()
	// otherExamples.RunCockroach10214()
	// otherExamples.RunEtcd5509()
	// otherExamples.RunEtcd6708()
	// otherExamples.RunGrpc3017()
	// otherExamples.RunKubernetes13135()
	// otherExamples.RunKubernetes62464()
	// otherExamples.RunKubernetes30872()
	// otherExamples.RunMoby4951()
	// otherExamples.RunMoby7559()
	// otherExamples.RunMoby17176()
	// otherExamples.RunMoby36114()
	// otherExamples.RunSyncthing4829()

	time.Sleep(time.Second * 3)

}
