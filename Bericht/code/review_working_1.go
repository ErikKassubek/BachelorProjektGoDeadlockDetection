func circularLocking() {
	var x Mutex
	var y Mutex
	ch := make(chan bool, 2)

	go func() {
		y.Lock()
		x.Lock()
		x.Unlock()
		y.Unlock()
		ch <- true
	}()
	go func() {
		x.Lock()
		y.Lock()
		y.Unlock()
		x.Unlock()
		ch <- true
	}()

	<-ch
	<-ch
}