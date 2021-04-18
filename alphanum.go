package validator

func IsValidAlphaNum(input string, length int) bool {

	if len(input) != length {
		return false
	}

	for _, c := range input {
		if c < 48 || c > 122 { // less than 0
			return false
		} else if c > 57 && c < 65 { // between 9 and A
			return false
		} else if c > 90 && c < 97 { // between Z and a
			return false
		}
	}
	return true
}
