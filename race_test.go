package main

import (
	"sync/atomic"
	"testing"
)

func TestDataRaceConditions(t *testing.T) {
	var state int32

	for i := 0; i < 10; i++ {
		go func(i int) {
			atomic.AddInt32(&state, int32(i))
		}(i)
	}

}
