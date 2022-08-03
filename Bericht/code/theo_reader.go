func readerLock() {
	var x RWMutex
	var y RWMutex
	var z RWMutex

	go func() {  // R1
		x.RLock()
		y.RLock()
		y.RUnLock()
		x.RUnLock()
	}

	go func() {  // R2
		y.RLock()
		x.RLock()
		x.RUnLock()
		y.RUnLock()
	}
}