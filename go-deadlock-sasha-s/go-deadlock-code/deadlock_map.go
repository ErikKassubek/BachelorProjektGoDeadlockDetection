//go:build go1.9
// +build go1.9

package main

import "sync"

// Map is sync.Map wrapper
type Map struct {
	sync.Map
}
