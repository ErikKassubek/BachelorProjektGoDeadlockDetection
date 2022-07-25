package measureRuntime

import (
	"fmt"
	"time"

	deadlock "github.com/ErikKassubek/Deadlock-Go"
)

// two routines, two locks per routine
func prog2x2() {
	x := deadlock.NewLock()
	y := deadlock.NewLock()
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
		y.Unlock()
		x.Unlock()

		ch <- true
	}()

	<-ch
	<-ch
}

// two routines, 1000 locks per routine
func prog2x100() {
	x := make([]*deadlock.Mutex, 100)
	for i := 0; i < 100; i++ {
		x[i] = deadlock.NewLock()
	}

	ch := make(chan bool, 2)
	ch2 := make(chan bool)

	go func() {
		<-ch2
		for i := 0; i < 100; i++ {
			x[i].Lock()
		}
		for i := 99; i >= 0; i-- {
			x[i].Unlock()
		}

		ch <- true
	}()

	go func() {
		for i := 99; i >= 0; i-- {
			x[i].Lock()
		}
		ch2 <- true
		for i := 0; i < 100; i++ {
			x[i].Unlock()
		}

		ch <- true
	}()

	<-ch
	<-ch
}

// 100 routines, 2 locks per routine
func prog100x2() {
	x := deadlock.NewLock()
	y := deadlock.NewLock()
	ch := make(chan bool, 100)

	for i := 0; i < 50; i++ {
		go func() {
			y.Lock()
			x.Lock()
			y.Unlock()
			x.Unlock()

			ch <- true
		}()
	}
	for i := 0; i < 50; i++ {
		<-ch
	}

	for i := 0; i < 50; i++ {
		go func() {
			x.Lock()
			y.Lock()
			y.Unlock()
			x.Unlock()

			ch <- true
		}()
	}
	for i := 0; i < 50; i++ {
		<-ch
	}
}

// 100 routines, 100 locks per routine
func prog100x100() {
	x := make([]*deadlock.Mutex, 100)
	for i := 0; i < 100; i++ {
		x[i] = deadlock.NewLock()
	}
	ch := make(chan bool, 100)

	for i := 0; i < 50; i++ {
		go func() {
			for i := 0; i < 100; i++ {
				x[i].Lock()
			}
			for i := 99; i >= 0; i-- {
				x[i].Unlock()
			}
			ch <- true
		}()
	}
	for i := 0; i < 50; i++ {
		<-ch
	}

	for i := 0; i < 50; i++ {
		go func() {
			for i := 99; i >= 0; i-- {
				x[i].Lock()
			}

			for i := 0; i < 100; i++ {
				x[i].Unlock()
			}

			ch <- true
		}()
	}
	for i := 0; i < 50; i++ {
		<-ch
	}
}

func RunTiming() {
	// sasha.Opts.OnPotentialDeadlock = func() {}
	// deadlock.SetCollectCallStack(false)
	// deadlock.SetCollectSingleLevelLockInformation(false)
	deadlock.SetPeriodicDetection(false)
	start := time.Now()

	prog100x100()

	durationRuntime := time.Since(start).Microseconds()

	deadlock.FindPotentialDeadlocks()

	durationTotal := time.Since(start).Microseconds()
	durationComprehensive := durationTotal - durationRuntime

	fmt.Println(durationRuntime)
	fmt.Println(durationComprehensive)
	fmt.Println(durationTotal)
}
