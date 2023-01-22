package main

import (
	"sync"
)

var (
	x    int64
	lock sync.Mutex
)
