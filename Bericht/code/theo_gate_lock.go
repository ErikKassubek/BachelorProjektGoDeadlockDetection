func gateLock() {
	var x Mutex
	var y Mutex
	var z Mutex

	go func() {  // R1
		z.Lock()
		y.Lock()
		x.Lock()
		x.Unlock()
		y.Unlock()
		z.Unlock()
	}

	go func() {  // R2
		z.Lock()
		x.Lock()
		y.Lock()
		y.Unlock()
		x.Unlock()
		z.Unlock()
	}
}