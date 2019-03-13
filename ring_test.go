package gencode

import (
	"fmt"
	"testing"
)

func TestNewRing(t *testing.T) {

	ring := NewRing(10, 100)
	for i := 0; i <= 500; i++ {
		fmt.Println(ring.Next())
	}
}
