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
	fmt.Println(d, ResolveTime(d))
}