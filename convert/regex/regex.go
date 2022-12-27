package regex

import "regexp"

// 匹配中国手机号码
func MatchChinaMobile(mobileNo string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNo)
}

// 匹配邮箱
func MatchEmail(email string) bool {
	//pattern := `\w+([-+.]\w+)@\w+([-.]\w+).\w+([-.]\w+)*` //匹配电子邮箱
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z].){1,4}[a-z]{2,4}$`

	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

// 匹配身份证号码
func MatchIdCard(idCard string) bool {
	regular := "^[1-9]\\d{5}(18|19|([23]\\d))\\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\\d{3}[0-9Xx]$"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(idCard)
}

// 匹配银行卡
func MatchBankCard(bankCard string) bool {
	regular := "^([1-9]{1})(\\d{14}|\\d{18})$"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(bankCard)
}
