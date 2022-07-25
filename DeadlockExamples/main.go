package main

import (
	"DeadlockExamples/selfWritten"

	deadlock "github.com/ErikKassubek/Deadlock-Go"
)

func main() {
	deadlock.SetMaxRoutines(3100)
	deadlock.SetPeriodicDetection(true)

	selfWritten.RunDeadlockGo()

	// selfWritten.RunSasha()

	// selfWritten.RunWithoutDetector()

	// otherExamples.RunCockroach584()
	// otherExamples.RunCockroach9935()
	// otherExamples.RunCockroach6181()
	// otherExamples.RunCockroach7504()
	// otherExamples.RunCockroach10214()
	// otherExamples.RunEtcd5509()
	// otherExamples.RunEtcd6708()
	// otherExamples.RunEtcd10492()
	// otherExamples.RunKubernetes13135()
	// otherExamples.RunKubernetes30872()
	// otherExamples.RunMoby4951()
	// otherExamples.RunMoby7559()
	// otherExamples.RunMoby17176()
	// otherExamples.RunMoby36114()
	// otherExamples.RunSyncthing4829()

	// time.Sleep(time.Second * 1)  // for otherExamples it is sometimes necessary to wait so that the detection is not started before the program has finished

	deadlock.FindPotentialDeadlocks()

	// measureRuntime.RunTiming()
}
