func cyclicLockingExample() {
	var v Mutex
	var w Mutex
	var x Mutex
	var y Mutex
	var z Mutex

	go func() { // R1
		v.Lock()
		w.Lock()
		w.Unlock()
		v.Unlock()
		y.Lock()
		z.Lock()
		z.Unlock()
		x.Lock()
		x.Unlock()
		y.Unlock()
	}()

	go func() { // R2
		w.Lock()
		x.Lock()
		x.Unlock()
		w.Unlock()
	}()

	go func() { // R3
		x.Lock()
		v.Lock()
		v.Unlock()
		x.Unlock()
	}
}