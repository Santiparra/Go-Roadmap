package main

func isValidPassword(password string) bool {
	length := len(password)
	hasUpper := false
	hasNumber := false

	for _, char := range password {
		if char >= 'A' && char <= 'Z' {
			hasUpper = true
		}
		if char >= '0' && char <= '9' {
			hasNumber = true
		}
	}

	return length >= 5 && length <= 12 && hasUpper && hasNumber
}