package main

import (
	"regexp"
)

const (
	phoneNumberRedactStub = "PII_PHONE_NUMBER"
)

var phoneNumberRegex *regexp.Regexp
var phoneNumberRegex2 *regexp.Regexp
var phoneNumberRegex3 *regexp.Regexp

func init() {
	// Regex source: https://github.com/solvvy/redact-pii/blob/main/src/built-ins/simple-regexp-patterns.ts
	// Modified the regex from that source to look for spaces too but to not
	// look for all numbers squished together because that caused false
	// positives with our customers' product IDs.

	phoneNumberRegex = regexp.MustCompile(`(\(?\+?[0-9]{1,2}\)?[-. ])?(\(?[0-9]{3}\)?|[0-9]{3})[-. ]([0-9]{3}[-. ][0-9]{4}\b)`)

	phoneNumberRegex2 = regexp.MustCompile(`\([0-9]{3}\) [0-9]{3}[-.][0-9]{4}`)

	// The exception to the rule that we don't match phone numbers when there is
	// no hyphen or period between each group is when there is "+1" preprended
	// to it. This gives us enough to look for without having false positive
	// matches.
	phoneNumberRegex3 = regexp.MustCompile(`\+1\d{10}`)
}

// redactPhoneNumber redacts credit card numbers from strings.
func redactPhoneNumber(src []byte) redactResult {
	numRedacted := int64(0)

	processed := phoneNumberRegex.ReplaceAllFunc(src, func(_ []byte) []byte {
		numRedacted += 1
		return []byte(phoneNumberRedactStub)
	})

	processed = phoneNumberRegex2.ReplaceAllFunc(processed, func(_ []byte) []byte {
		numRedacted += 1
		return []byte(phoneNumberRedactStub)
	})

	processed = phoneNumberRegex3.ReplaceAllFunc(processed, func(_ []byte) []byte {
		numRedacted += 1
		return []byte(phoneNumberRedactStub)
	})

	return redactResult{
		numRedacted: numRedacted,
		redacted:    processed,
	}
}
