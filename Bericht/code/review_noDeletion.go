func noDeletion() {
	var x Mutex
	var y Mutex

	x.Lock()
	y.Lock()
	y.Unlock()
	x.Unlock()

	y.Lock()
	x.Lock()
	x.Unlock()
	y.Unlock()
}