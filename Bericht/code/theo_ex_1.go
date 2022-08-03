func cyclicLockingExample() {
	var x Mutex
	var y Mutex

	go func() { // R1
		y.Lock()
		x.Lock()
		x.Unlock()
		y.Unlock()
	}()

	go func() { // R2
		x.Lock()
		y.Lock()
		y.Unlock()
		x.Unlock()
	}()
}