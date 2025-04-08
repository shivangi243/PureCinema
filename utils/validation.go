package utils

import "regexp"

func IsPasswordStrong(pw string) bool {
	if len(pw) < 8 {
		return false
	}

	hasNumber, _ := regexp.MatchString(`[0-9]`, pw)
	hasSpecial, _ := regexp.MatchString(`[!@#~$%^&*()_+{}\[\]:;"'<>,.?\\|]`, pw)

	return hasNumber && hasSpecial
}
