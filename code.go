package gencode

import (
	"fmt"
	"strconv"
	"time"
)

// 生成10位时间码
func GenTime() int64 {
	tm := time.Now()
	tmStart, _ := time.ParseInLocation("2006", strconv.Itoa(tm.Year()), time.Local)
	//一年的秒数是最大8位
	seconds := tm.Sub(tmStart).Nanoseconds() / 1e9
	//两位的年
	year := tm.Year() % 100
	return int64(year*1e8) + seconds
}

// 校验时间码
func ResolveTime(t int64) time.Time {
	// 获取当前年份，比如19年是19
	year := int(t/1e8)
	tmStart, _ := time.ParseInLocation("2006", "20"+ToStr(year), time.Local)
	// 反算秒数
	seconds := t - int64(year*1e8)
	return tmStart.Add(time.Duration(seconds) * time.Second)
}

// 生成唯一码
// rn 是从散列中取的随机码
func GenCode(rn string) string {
	// 10 位的时间信息
	t := GenTime()
	// 混淆
	mt := MixCode(ToStr(t), rn)
	code := mt + rn
	// 加校验位
	return mt + rn + ToStr(LuhnGenerate(StrTo(code).MustInt64()))
}

func ResolveCode(code string) (*time.Time, error) {
	//得到前9位时间码转换为时间对象
	mt := code[:10]
	rn := code[10:15]

	fmt.Println("rcode",mt,rn)
	t := DeMixCode(mt, rn)

	fmt.Println("t",t)

	tn, err := StrTo(t).Int64()

	if err != nil {
		return nil, err
	}

	tm := ResolveTime(tn)

	return &tm, nil
}
