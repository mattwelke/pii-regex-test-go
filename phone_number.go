package main

import (
	"regexp"
)

const (
	phoneNumberRedactStub = "PII_PHONE_NUMBER"
)

var phoneNumberRegex *regexp.Regexp

func init() {
	// Regex source:
	phoneNumberRegex = regexp.MustCompile(`(\(?\+?[0-9]{1,2}\)?[-. ]?)?(\(?[0-9]{3}\)?|[0-9]{3})[-. ]?([0-9]{3}[-. ]?[0-9]{4}|\b[A-Z0-9]{7}\b)`)
}

// redactPhoneNumber redacts credit card numbers from strings.
func redactPhoneNumber(src []byte) redactResult {
	numRedacted := int64(0)

	processed := phoneNumberRegex.ReplaceAllFunc(src, func(_ []byte) []byte {
		numRedacted += 1
		return []byte(phoneNumberRedactStub)
	})

	return redactResult{
		numRedacted: numRedacted,
		redacted:    processed,
	}
}
