package gencode

import (
	"fmt"
	"testing"

	"github.com/ilibs/gencode/internal"
)

func TestLuhnValid(t *testing.T) {
	cases := []int64{
		123456789,
		29396257166338,
		59385787196327,
		87894815624836,
		81261619028203,
		2407330249449,
		22977250769919,
		24640552461682,
		26871854663813,
		38593266375535,
		46569274383501,
	}

	for _, i := range cases {
		n := LuhnGenerate(i)
		m := fmt.Sprintf("%d%d", i, n)
		fmt.Println(m)
		if ok := LuhnValid(internal.StrTo(m).MustInt64()); !ok {
			t.Fatalf("Valid(%d): %d\n\t Got: %t", i, n, ok)
		}
	}
}

func TestLuhnGenerate(t *testing.T) {
	fmt.Println("Luhn Number:", LuhnGenerate(123456789))
	if ok := LuhnValid(1234567897); !ok {
		t.Fatalf("Valid Error")
	}
}
