import "github.com/ErikKassubek/Deadlock-Go"

func main() {
	x := deadlock.NewLock()
	y := deadlock.NewLock()

	// make sure, that the program does not terminate
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

	// start the final deadlock detection
	deadlock.FindPotentialDeadlocks()
}
