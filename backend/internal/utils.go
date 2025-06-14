package internal

import (
	"regexp"
	"unicode"

	"github.com/amha-mersha/sanqa-suq/internal"
	"golang.org/x/crypto/bcrypt"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

func ValidateEmail(email string) bool {
	return emailRegex.MatchString(email)
}

func ValidatePassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	var hasUpper, hasLower, hasNumber, hasSpecial bool

	for _, ch := range password {
		switch {
		case unicode.IsUpper(ch):
			hasUpper = true
		case unicode.IsLower(ch):
			hasLower = true
		case unicode.IsDigit(ch):
			hasNumber = true
		case unicode.IsPunct(ch) || unicode.IsSymbol(ch):
			hasSpecial = true
		}
	}

	return hasUpper && hasLower && hasNumber && hasSpecial
}

var phoneRegex = regexp.MustCompile(`^\+?[0-9]{10,15}$`)

func ValidatePhoneNumber(phone string) bool {
	return phoneRegex.MatchString(phone)
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", internal.InternalError("FAILED_TO_HASH_PASSWORD", err)
	}
	return string(hashedPassword), nil
}
