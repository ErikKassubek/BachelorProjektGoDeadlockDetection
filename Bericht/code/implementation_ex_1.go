import (
	deadlock "github.com/ErikKassubek/Deadlock-Go"
)

func potentialDeadlock() {
	x := deadlock.NewLock()
	y := deadlock.NewLock()
	ch := make(chan bool, 2)

	go func() {
		deadlock.NewRoutine()

		x.Lock()
		y.Lock()
		time.Sleep(time.Second)
		y.Unlock()
		x.Unlock()

		ch <- true
	}()

	go func() {
		deadlock.NewRoutine()

		y.Lock()
		x.Lock()
		time.Sleep(time.Second)
		x.Unlock()
		y.Unlock()

		ch <- true
	}()

	<-ch
	<-ch
}

func main() {
	deadlock.Initalize()
	defer deadlock.FindPotentialDeadlocks()
	potentialDeadlock()
}