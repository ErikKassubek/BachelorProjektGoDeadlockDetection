func nestedGoRoutines() {
	var x deadlock.Mutex
	var y deadlock.Mutex
	ch := make(chan bool)

	go func() {
		y.Lock()
		// nested routine
		go func() {
			x.Lock()
			x.Unlock()
			ch <- true
		}()
		<-ch
		y.Unlock()
	}()

	go func() {
		x.Lock()
		y.Lock()
		y.Lock()
		x.Lock()
	}()
}