package gencode

import (
	"fmt"
	"testing"

	"github.com/ilibs/gencode/internal"
)

func TestMixCode(t *testing.T) {
	key := "2018091907"

	for i := 100000; i < 100999; i++ {
		//fmt.Println("before", before)
		before := internal.ToStr(i)
		s := MixCode(before, key)
		//fmt.Println("end   ", s)
		s2 := DeMixCode(s, key)
		//fmt.Println("before", s2)

		if before != s2 {
			t.Errorf("mix erro before %s but get %s", before, s2)
		}
	}
}

func TestMixRepeat2(t *testing.T) {
	key := "20180919"
	m := make(map[string]bool)

	for i := 1000000000; i < 1000000999; i++ {
		//fmt.Println("before", before)
		before := internal.ToStr(i)
		s := MixCode(before, key)
		if i < 1000000010 {
			fmt.Println(s)
		}

		if m[s] {
			t.Fatal("重复了", s)
		}
		m[s] = true

	}
}