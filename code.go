package gencode

import (
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
func ResolveTime(t int64) (*time.Time, error) {
	// 获取当前年份，比如19年是19
	year := int(t / 1e8)
	tmStart, err := time.ParseInLocation("2006", "20"+ToStr(year), time.Local)
	if err != nil {
		return nil, err
	}

	// 反算秒数
	seconds := t - int64(year*1e8)
	tt := tmStart.Add(time.Duration(seconds) * time.Second)
	return &tt, nil
}

// 生成唯一码
// rn 是从散列中取的随机码
func GenCode(rn string) string {
	// 10 位的时间信息
	t := GenTime()
	// 必须保证rn在前面，因为混淆之后第一位数可能是0，会导致后续求校验码错误
	return rn + ToStr(t)
}

func ResolveCode(code string) (*time.Time, error) {
	//得到前10位时间码转换为时间对象
	t := code[5:15]
	tn, err := StrTo(t).Int64()

	if err != nil {
		return nil, err
	}

	tm, err := ResolveTime(tn)
	if err != nil {
		return nil, err
	}

	return tm, nil
}
