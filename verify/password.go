package verify

import (
	"regexp"
	"unicode"
)

func VerifyPasswordFormat(password string) bool {
	var hasNumber, hasLetter, hasUpperCase, hasLowercase, hasSpecial bool
	for _, c := range password {
		switch {
		case unicode.IsNumber(c):
			hasNumber = true //是否包含数字
		case unicode.IsLetter(c):
			hasLetter = true //是否包含字母
			if unicode.IsUpper(c) {
				hasUpperCase = true //是否包含大写字母
			}
			if unicode.IsLower(c) {
				hasLowercase = true //是否包含小写字母
			}
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			hasSpecial = true //是否包含特殊符号
		}

	}
	//不检查特殊字符
	hasSpecial = true

	//不检查大写
	hasUpperCase = true

	//不检查小写
	hasLowercase = true

	return hasNumber && hasLetter && hasUpperCase && hasLowercase && hasSpecial && len(password) > 7 && len(password) < 21
}

func VerifyPayPassFormat(password string) bool {
	regular := "^[0-9]{6,6}$"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(password)

}
