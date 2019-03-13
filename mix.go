package gencode

import (
	"strconv"
)

// 混淆算法
// org 使用 key 混淆
func MixCode(org string, key string) string {
	lens1 := len(org)
	lens2 := len(key)
	s := ""
	for i := 0; i < lens1; i++ {
		k := i % lens2
		s += digitMod(org[i], key[k])
	}
	return s
}

// 反混淆算法
// sec 使用 key 反混淆
func DeMixCode(sec string, key string) string {
	lens1 := len(sec)
	lens2 := len(key)
	s := ""
	for i := 0; i < lens1; i++ {
		k := i % lens2
		s += deDigitMod(sec[i], key[k])
	}
	return s
}

// 0 == 48
func digitMod(s1 uint8, s2 uint8) string {
	n := int((s1-48)+(s2-48)) % 10
	return strconv.Itoa(int(n))
}

func deDigitMod(s1 uint8, s2 uint8) string {
	ds1 := s1 - 48
	ds2 := s2 - 48
	var n uint8
	if ds2 > ds1 {
		n = ds2 - ds1
	} else {
		n = 10 - (ds1 - ds2)
	}
	n = 10 - n
	return strconv.Itoa(int(n))
}
