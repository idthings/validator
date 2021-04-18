package validator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInvalidEmailAddress(t *testing.T) {

	invalidEmailAddresses := [...]string{
		"",
		"missingatsign",
		"@missing.user",
		"missing@fqdn",
		"user@more.than.four.domain.segments",
		"user@invalidchar+.indomain",
	}

	for _, email := range invalidEmailAddresses {
		assert.Equal(t, false, IsValidEmail(email), email)
	}
}

func TestValidEmailAddress(t *testing.T) {

	validEmailAddresses := [...]string{
		"user@domain.com",
		"underscore_in_name@domain.com",
	}

	for _, email := range validEmailAddresses {
		assert.Equal(t, true, IsValidEmail(email), email)
	}
}
