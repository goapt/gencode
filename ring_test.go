package gencode

import (
	"fmt"
	"testing"
	"time"
)

func TestNewRing(t *testing.T) {

	ring := NewRing(10, 1000, time.Second)
	for i := 0; i <= 500; i++ {
		fmt.Println(ring.Next())
	}
}
