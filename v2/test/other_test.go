package test

import (
	"log"
	"sync/atomic"
	"testing"
)

func TestCAS(t *testing.T) {
	var i int32 = 5
	for atomic.CompareAndSwapInt32(&i, i, i+1) {
		break
	}
	log.Println(i)
}
