package gencode

import (
	"errors"
	"time"
)

type Manager struct {
	Prefix string //前缀 可为空
	Key    string //混淆码
	Ring   *Ring  //散列迭代器
}

// 生成实例
func New(prefix string, key string) *Manager {
	m := &Manager{
		Prefix: prefix,
		Key:    key,
		Ring:   NewRing(10000, 99999),
	}
	return m
}

// 生成码
func (m *Manager) Get() string {
	return m.Prefix + MixCode(GenCode(ToStr(m.Ring.Next())),m.Key)
}

// 校验码
func (m *Manager) Verify(s string) (*time.Time, error) {
	if len(s) != 16 {
		return nil, errors.New("code parse error: length must be 16")
	}

	//去除前缀，得到码
	code := s[len(m.Prefix):]
	code = DeMixCode(code,m.Key)

	//是否为纯数字
	coden, err := StrTo(code).Int64()
	if err != nil {
		return nil, err
	}

	//检查校验位
	if !LuhnValid(coden) {
		return nil, errors.New("code check error")
	}

	return ResolveCode(code)
}
