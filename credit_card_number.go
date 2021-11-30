package main

import (
	"regexp"
)

const (
	creditCardRedactStub = "PII_CREDIT_CARD"
)

var creditCardRegex *regexp.Regexp

func init() {
	creditCardRegex = regexp.MustCompile(`\d{4}[ -]?\d{4}[ -]?\d{4}[ -]?\d{4}|\d{4}[ -]?\d{6}[ -]?\d{4}\d?`)
}

// redactCreditCard redacts credit card numbers from strings.
func redactCreditCard(src []byte) redactResult {
	numRedacted := int64(0)

	processed := creditCardRegex.ReplaceAllFunc(src, func(_ []byte) []byte {
		numRedacted += 1
		return []byte(creditCardRedactStub)
	})

	return redactResult{
		numRedacted: numRedacted,
		redacted:    processed,
	}
}
