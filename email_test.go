package validator

import (
	"regexp"
	"strings"
)

func IsValidEmail(input string) bool {
	if len(input) < 3 || len(input) > 254 {
		return false
	}

	emailSegments := strings.Split(input, "@")
	if len(emailSegments) < 2 {
		return false // no @ in input
	}

	var emailUserRegex = regexp.MustCompile("^[a-zA-Z0-9.-]+$")

	userSegment := emailSegments[0]
	if !emailUserRegex.MatchString(userSegment) {
		return false
	}

	var emailDomainRegex = regexp.MustCompile("^[a-zA-Z0-9.-]+$")

	domainSegments := strings.Split(emailSegments[1], ".")
	if len(domainSegments) < 2 || len(domainSegments) > 4 {
		return false
	}

	for _, segment := range domainSegments {
		if !emailDomainRegex.MatchString(segment) {
			return false
		}
	}
	return true
}
