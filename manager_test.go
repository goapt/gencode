package gencode

import (
	"fmt"
	"testing"
	"time"
)

var manager = New("", "20180919")

func TestResolveCode(t *testing.T) {

	code := manager.Get()
	fmt.Println(code)
	c, err := manager.Verify(code)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(c, err)
}

func TestGenCodeRepeat(t *testing.T) {
	m := make(map[string]bool)
	for n := 1; n < 100; n++ {
		for i := 10000; i < 99999; i++ {
			code := manager.Get()
			if i < 10 {
				fmt.Println(code)
			}

			if m[code] {
				t.Fatal("重复了", code)
			}
			m[code] = true
		}
		time.Sleep(1 * time.Second)
	}
}
