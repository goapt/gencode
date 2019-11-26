package gencode

import (
	"errors"
	"time"

	"github.com/goapt/gencode/internal"
)

type Manager struct {
	Prefix string //前缀 可为空
	Key    string //混淆码
	Ring   *Ring  //散列迭代器
	Luhn   bool   //是否在末尾添加校验码
}

// 生成实例
func New(prefix string, key string, luhn bool) *Manager {
	m := &Manager{
		Prefix: prefix,
		Key:    key,
		Ring:   NewRing(10000, 99999, time.Second),
		Luhn:   luhn,
	}
	return m
}

// 生成码
func (m *Manager) Get() string {
	//得到15位编码
	code := GenCode(internal.ToStr(m.Ring.Next()))

	if m.Luhn {
		code = code + internal.ToStr(LuhnGenerate(internal.StrTo(code).MustInt64()))
	}
	//前缀 + 混淆
	return m.Prefix + MixCode(code, m.Key)
}

// 校验码
func (m *Manager) Verify(s string) (*time.Time, error) {
	//去除前缀，得到码
	code := s[len(m.Prefix):]
	code = DeMixCode(code, m.Key)

	//是否为纯数字
	coden, err := internal.StrTo(code).Int64()
	if err != nil {
		return nil, err
	}

	//检查校验位
	if m.Luhn && !LuhnValid(coden) {
		return nil, errors.New("code check error")
	}

	return ResolveCode(code)
}
