package ytime

import "strings"

type DateTime struct {
	unix  int64 //秒级时间戳
	year  int
	month int
	day   int
	hour  int
	min   int
	sec   int
	zone  TimeZone
}

const DefaultFormat = "%Y-%m-%d %H:%M:%S"

func (d DateTime) UnixMS() int64 {
	return d.unix
}

// ToDefaultStr 默认格式化
func (d DateTime) ToDefaultStr() string {
	return d.Format(DefaultFormat)
}

func (d DateTime) Format(formatStr string) string {
	var dateStr strings.Builder
	length := len(formatStr)
	for i := 0; i < length; {
		char := formatStr[i]
		if char == '%' {
			if i+1 >= length {
				break
			}
			val := formatStr[i+1]
			switch val {
			case 'Y':
				dateStr.WriteRune(rune(d.year))
			case 'y':
			case 'm':
				dateStr.WriteRune(rune(d.month))
			}
		} else {
			dateStr.WriteByte(char)
			i++
		}
	}
	return ""
}
