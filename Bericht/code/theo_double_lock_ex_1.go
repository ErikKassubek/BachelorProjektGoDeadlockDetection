func doubleLockingExample() {
	var x Mutex

	go func() { // R1
		x.Lock()
		x.Lock()
		x.Unlock()
	}()
}