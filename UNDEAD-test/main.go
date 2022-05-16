package main

import undead "github.com/ErikKassubek/UNDEAD-go"

func test1() {
	var x undead.Mutex
	var y undead.Mutex
	ch := make(chan bool, 2)

	go func() {
		r := undead.NewRoutine()
		x.Lock(r)
		y.Lock(r)
		y.Unlock()
		x.Unlock()
		ch <- true
	}()

	go func() {
		r := undead.NewRoutine()
		y.Lock(r)
		x.Lock(r)
		x.Unlock()
		y.Unlock()
		ch <- true
	}()

	<-ch
	<-ch
}

func main() {
	test1()
}
