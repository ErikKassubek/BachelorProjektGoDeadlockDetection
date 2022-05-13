func guardLocks() {
	var x deadlock.Mutex
	var y deadlock.Mutex
	var z deadlock.Mutex

	ch := make(chan bool, 2)

	go func() {
		z.Lock()
		y.Lock()
		x.Lock()
		x.Unlock()
		y.Unlock()
		z.Unlock()
		ch <- true
	}()
	go func() {
		z.Lock()
		x.Lock()
		y.Lock()
		y.Unlock()
		x.Unlock()
		z.Unlock()
		ch <- true
	}()

	<-ch
	<-ch
}