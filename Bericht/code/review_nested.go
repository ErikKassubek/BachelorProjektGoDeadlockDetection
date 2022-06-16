func nestedRoutines() {
	x := deadlock.NewLock()
	y := deadlock.NewLock()
	ch := make(chan bool)
	ch2 := make(chan bool)

	go func() {
		x.Lock()
		go func() {
			y.Lock()
			y.Unlock()
			ch <- true
		}()
		<-ch
		x.Unlock()
		ch2 <- true
	}()
	go func() {
		y.Lock()
		x.Lock()
		x.Unlock()
		y.Unlock()
		ch2 <- true
	}()

	<-ch2
	<-ch2
}