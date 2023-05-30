package singleton

import (
	"testing"
	"time"
)

func TestSingleton(t *testing.T) {
	for i := 0; i < 100; i++ {
		go func() {
			New()
		}()
	}
	time.Sleep(time.Second * 2)
}
