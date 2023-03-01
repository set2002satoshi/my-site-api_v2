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
