package convert

import (
	"regexp"
)

// 用*替换手机号码的中间4位
func ReplaceByStar(mobile string) string {
	var re = regexp.MustCompile(`(\d{3})\d{4}(\d{4})`)
	s := re.ReplaceAllString(mobile, `$1****$2`)

	return s
}

//替换姓名的后两位
func ReplaceNameWithAsterisk(name string) string {
	n := len(name)
	str := ""
	r := []rune(name)
	if n == 6 {
		str += string(r[:1]) + "*"
		return str
	}
	if n == 9 {
		str += string(r[:1]) + "*" + string(r[len(r)-1])
		return str
	}
	if n == 12 {
		str += string(r[:1]) + "**" + string(r[len(r)-1])
		return str
	}
	return ""
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func Substr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}
	return string(rs[start:end])
}

// 星号替换身份证号码
func ReplaceIdCardByStar(idCard string) string {

	// 18位身份证 ^(\d{17})([0-9]|X)$
	// 匹配规则:
	// 1. 17位数字 2. 18位数字或者X
	// (^\d{15}$) 15位身份证
	// (^\d{18}$) 18位身份证
	// (^\d{17}(\d|X|x)$) 18位身份证 最后一位为X的用户

	regRuler := "(^\\d{15}$)|(^\\d{18}$)|(^\\d{17}(\\d|X|x)$)"

	var re = regexp.MustCompile(regRuler)

	if re.MatchString(idCard) {
		if len(idCard) == 15 {
			return idCard[:4] + "*******" + idCard[14:]
		} else {
			return idCard[:4] + "**********" + idCard[14:]
		}
	}

	return idCard
}

// 星号替换邮箱
func ReplaceEmailByStar(email string) string {
	var re = regexp.MustCompile(`(\w{1,3})\w+@(\w+)(\.\w+)`)
	s := re.ReplaceAllString(email, `$1***@$2$3`)

	return s
}
