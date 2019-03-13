package gencode

import (
	"fmt"
	"testing"
)

func TestMixCode(t *testing.T) {
	key := "2018091907"

	for i :=100000; i < 999999;i++ {
		//fmt.Println("before", before)
		before := ToStr(i)
		s := MixCode(before, key)
		//fmt.Println("end   ", s)
		s2 := DeMixCode(s, key)
		//fmt.Println("before", s2)

		if before != s2 {
			t.Errorf("mix erro before %s but get %s",before,s2)
		}
	}
}



func TestMixRepeat(t *testing.T) {
	ring := NewRing(10000, 99999)
	m := make(map[string]bool)

	for i :=1000000000; i < 1000999999;i++ {
		//fmt.Println("before", before)
		before := ToStr(i)
		rn := ToStr(ring.Next())
		s := MixCode(before, ToStr(ring.Next()))
		if i < 1000000010 {
			fmt.Println(s+rn)
		}

		if m[s+rn] {
			t.Fatal("重复了", s)
		}
		m[s+rn] = true

	}
}
