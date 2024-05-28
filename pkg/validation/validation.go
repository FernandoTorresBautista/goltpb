package validation

import "regexp"

// IsNumeric to validate the phone number
func IsNumeric(input string) bool {
	regex := regexp.MustCompile("^[0-9]+$")
	return regex.MatchString(input)
}

// ValidatePassword check the requirements for the password
func ValidatePassword(password string) bool {
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasDigit := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[@$&]`).MatchString(password)
	lengthValid := len(password) >= 6 && len(password) <= 12
	return lengthValid && hasLower && hasUpper && hasDigit && hasSpecial
}

// ValidateEmail check the format of the email
func ValidateEmail(email string) bool {
	regex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return regex.MatchString(email)
}
