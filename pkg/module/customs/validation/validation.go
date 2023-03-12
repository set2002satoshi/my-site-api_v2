package validation

import "regexp"

const (
	_mailAddressPattern = `^[a-zA-Z0-9_+-]+(\.[a-zA-Z0-9_+-]+)*@([a-zA-Z0-9][a-zA-Z0-9-]*[a-zA-Z0-9]*\.)+[a-zA-Z]{2,}$`
)

// 正しくない時はfalseを返す
func ValidEmail(email string) bool {
	r := regexp.MustCompile(_mailAddressPattern)
	return r.MatchString(email)
}

// 正しくない時にtrueを返す。日本語の場合５文字以上
func IsWhitespaceOrLessThan16Characters(s string) bool {
	return len(s) <= 16
}