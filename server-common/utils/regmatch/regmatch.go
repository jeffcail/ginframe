package regmatch

import (
	"fmt"
	"regexp"
)

const (
	PHONE            = `[1](([3][0-9])|([4][5-9])|([5][0-3,5-9])|([6][5,6])|([7][0-8])|([8][0-9])|([9][1,8,9]))[0-9]{8}`
	LINDLINENUMBER   = `\b(0\d{2,3}-\d{7,8}|\(?0\d{2,3}[)-]?\d{7,8}|\(?0\d{2,3}[)-]*\d{7,8})\b`
	RECORDNO         = `[1-9]\d{5}(18|19|20|(3\d))\d{2}((0[1-9])|(1[0-2]))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]`
	PASSPORTNUMBER   = `([a-zA-z]|[0-9]){5,17}`
	PASSPORT         = `[HMhm]{1}([0-9]{10}|[0-9]{8})`
	IPV4             = `(((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(.|$)){4})`
	IPV6             = `("["([0-9a-fA-F]+:){7}[0-9a-fA-F]+"]")|("["0x[0-9a-fA-F]+([0-9a-fA-F]+|:)*"::"([0-9a-fA-F]+|:)*"]")|("["([0-9a-fA-F]+|:)*"::"([0-9a-fA-F]+|:)*"]")|("["([0-9a-fA-F]+|:)*"::"([0-9a-fA-F]+|:)*([0-9]"."){3}[0-9]"]")`
	MAC              = `(?:(?:(?:[a-f0-9A-F]{2}:){5})|(?:(?:[a-f0-9A-F]{2}-){5}))[a-f0-9A-F]{2}`
	EMAIL            = `([\w.\_]{2,10})@(\w{1,}).([a-z]{2,4})`
	SOCIALCREDITCODE = `[\dANY]{1}\d{7}[0-9A-HJ-NPQRTUWXY]{10}`
	URL              = `^(((ht|f)tps?):\/\/)?([^!@#$%^&*?.\s-]([^!@#$%^&*?.\s]{0,63}[^!@#$%^&*?.\s])?\.)+[a-z]{2,6}\/?`
	URLPORT          = `^((ht|f)tps?:\/\/)?[\w-]+(\.[\w-]+)+:\d{1,5}\/?$`
	AMOUNT           = `(?:^[1-9]([0-9]+)?(?:\.[0-9]{1,2})?$)|(?:^(?:0)$)|(?:^[0-9]\.[0-9](?:[0-9])?$)`
)

// RegMatchPhone reg match phone
func RegMatchPhone(phone string) bool {
	matched, err := regexp.MatchString(PHONE, phone)
	if err != nil || !matched {
		return false
	}
	return true
}

// RegLandLineNumber 座机号 判定数据内容是否包含中国固定电话号码
func RegLandLineNumber(lindLineNumber string) bool {
	matched, err := regexp.MatchString(LINDLINENUMBER, lindLineNumber)
	if err != nil || !matched {
		return false
	}
	return true
}

// RegIDCard 18位身份证号
func RegIDCard(idCard string) bool {
	matched, err := regexp.MatchString(RECORDNO, idCard)
	if err != nil || !matched {
		return false
	}
	return true
}

// PassportNumber 护照编号
func PassportNumber(passportNumber string) bool {
	matched, err := regexp.MatchString(RECORDNO, passportNumber)
	if err != nil || !matched {
		return false
	}
	return true
}

// Passport 港澳通行证
func Passport(passport string) bool {
	matched, err := regexp.MatchString(PASSPORT, passport)
	if err != nil || !matched {
		return false
	}
	return true
}

// RegIPV4 判定数据内容是否包含IP地址
func RegIPV4(ipv4 string) bool {
	matched, err := regexp.MatchString(IPV4, ipv4)
	if err != nil || !matched {
		return false
	}
	return true
}

// RegIPV6 判定数据内容是否包含IPV6地址
func RegIPV6(ipv6 string) bool {
	matched, err := regexp.MatchString(IPV6, ipv6)
	if err != nil || !matched {
		return false
	}
	return true
}

// RegMAC 判定数据内容是否包含IP地址
func RegMAC(mac string) bool {
	matched, err := regexp.MatchString(MAC, mac)
	if err != nil || !matched {
		return false
	}
	return true
}

// RegEmail 电子邮箱地址
func RegEmail(email string) bool {
	matched, err := regexp.MatchString(EMAIL, email)
	if err != nil || !matched {
		return false
	}
	return true
}

// RegSocialCreditCode 统一社会信用代码
func RegSocialCreditCode(socialCreditCode string) bool {
	matched, err := regexp.MatchString(SOCIALCREDITCODE, socialCreditCode)
	if err != nil || !matched {
		return false
	}
	return true
}

// CheckPasswordLever 密码常用正则表达式 密码强度必须为字⺟⼤⼩写+数字+符号，9位以上
func CheckPasswordLever(ps string) error {
	if len(ps) < 9 {
		return fmt.Errorf("password len is < 9")
	}
	num := `[0-9]{1}`
	a_z := `[a-z]{1}`
	A_Z := `[A-Z]{1}`
	symbol := `[!@#~$%^&*()+|_]{1}`
	if b, err := regexp.MatchString(num, ps); !b || err != nil {
		return fmt.Errorf("password need num :%v", err)
	}
	if b, err := regexp.MatchString(a_z, ps); !b || err != nil {
		return fmt.Errorf("password need a_z :%v", err)
	}
	if b, err := regexp.MatchString(A_Z, ps); !b || err != nil {
		return fmt.Errorf("password need A_Z :%v", err)
	}
	if b, err := regexp.MatchString(symbol, ps); !b || err != nil {
		return fmt.Errorf("password need symbol :%v", err)
	}
	return nil
}

// RegUrlPort 网址URL（带端口号，如：https://www.baidu.com:8080/）
func RegUrlPort(urlPort string) bool {
	matched, err := regexp.MatchString(URLPORT, urlPort)
	if err != nil || !matched {
		return false
	}
	return true
}

// RegUrl 网址URL（不带端口号，如：https://www.baidu.com/）
func RegUrl(url string) bool {
	matched, err := regexp.MatchString(URL, url)
	if err != nil || !matched {
		return false
	}
	return true
}

// RegAmount 金额（正数，可有最多两位小数，如：8.99）
func RegAmount(amount string) bool {
	matched, err := regexp.MatchString(AMOUNT, amount)
	if err != nil || !matched {
		return false
	}
	return true
}
