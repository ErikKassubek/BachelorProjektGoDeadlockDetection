func falsePositive() {
	var x deadlock.Mutex
	finished := make(chan bool)

	go func() {
		// first go routine
		x.Lock()
		time.Sleep(40 * time.Second)
		x.Unlock()
	}()

	go func() {
		// second go routine
		time.Sleep(2 * time.Second)
		x.Lock()
		x.Unlock()
		finished <- true
	}()

	<-finished
}