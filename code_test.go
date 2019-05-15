package gencode

import (
	"fmt"
	"testing"
)

func TestGenTime(t *testing.T) {
	fmt.Println(GenTime())
}

func TestResolveTime(t *testing.T) {
	d := GenTime()
	mt, err := ResolveTime(d)
	fmt.Println(d, mt, err)
}
