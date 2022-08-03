func periodicalDetection() {
	var w Mutex 
	var x Mutex
	var y Mutex
	var z Mutex
	
	go func() { // R1
		w.Lock()
		x.Lock()
		x.UnLock()
		w.UnLock()
		y.Lock()
		z.Lock()
		z.UnLock()
		y.UnLock()
		// periodical detection is started here
	}

	go func() { // R2
		x.Lock()
		w.Lock()
		w.UnLock()
		x.UnLock()
		// periodical detection is started here
		z.Lock()
		y.Lock()
		y.UnLock()
		z.UnLock()
	}
}