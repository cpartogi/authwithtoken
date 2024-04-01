package usecase

import (
	"authwithtoken/domain/auth/model"
	"authwithtoken/lib/constant"
	"strings"
	"unicode"
)

func isDataValid(data model.Users) (invalidMessages []string, isValid bool) {

	if len(data.FullName) < 3 || len(data.FullName) > 60 {
		invalidMessages = append(invalidMessages, constant.FullNameCharLength)
	}

	if len(data.PhoneNumber) < 10 || len(data.PhoneNumber) > 13 {
		invalidMessages = append(invalidMessages, constant.PhoneNumberCharLength)
	}

	if !strings.HasPrefix(data.PhoneNumber, "+62") {
		invalidMessages = append(invalidMessages, constant.PhoneNumberIndonesian)
	}

	if len(data.Password) < 6 || len(data.Password) > 64 {
		invalidMessages = append(invalidMessages, constant.PassWordCharLength)
	}

	if len(data.Email) == 0 {
		invalidMessages = append(invalidMessages, constant.EmailRequired)
	}

	if !isValidPasswordChar(data.Password) {
		invalidMessages = append(invalidMessages, constant.PasswordReqChar)
	}

	if len(invalidMessages) > 0 {
		return invalidMessages, false
	} else {
		return []string{""}, true
	}
}

func isValidPasswordChar(s string) bool {
	var hasUpperCase, hasNumber, hasSpecial bool

	for _, char := range s {
		switch {
		case unicode.IsUpper(char):
			hasUpperCase = true
		case unicode.IsNumber(char):
			hasNumber = true
		case !unicode.IsLetter(char) && !unicode.IsNumber(char):
			hasSpecial = true
		}

		if hasUpperCase && hasNumber && hasSpecial {
			return true
		}
	}

	return false
}
