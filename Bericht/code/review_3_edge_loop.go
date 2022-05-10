func threeEdgeLoop() {
	var x Mutex
	var y Mutex
	var z Mutex
	ch := make(chan bool, 3)

	go func() {
		// first routine
		x.Lock()
		y.Lock()
		y.Unlock()
		x.Unlock()
		ch <- true
	}()

	go func() {
		// second routine
		y.Lock()
		z.Lock()
		z.Unlock()
		y.Unlock()
		ch <- true
	}()

	go func() {
		// third routine
		z.Lock()
		x.Lock()
		x.Unlock()
		z.Unlock()
		ch <- true
	}()

	<-ch
	<-ch
	<-ch
}