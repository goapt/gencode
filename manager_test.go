package gencode

import (
	"fmt"
	"testing"
)

var manager = New("1", "20180919", false)

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
	for i := 1; i < 1000000; i++ {
		code := manager.Get()
		if i < 100 {
			fmt.Println(code)
		}

		_, err := manager.Verify(code)
		if err != nil {
			t.Fatal("验证失败", code, err)
		}

		if m[code] {
			t.Fatal("重复了", code)
		}
		m[code] = true
	}
}
