import "github.com/ErikKassubek/Deadlock-Go"

func main() {
	defer deadlock.FindPotentialDeadlocks()
	x := deadlock.NewLock()
	y := deadlock.NewLock()

	// make sure, that program does not terminates
	// before all routines have terminated
	ch := make(chan bool, 2)

	go func() {
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
}
