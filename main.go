package main

import (
	"fmt"
	"regexp"
)

const (
	creditCardRedactStub = "PII_CREDIT_CARD"
)

var creditCardRegex *regexp.Regexp

func init() {
	creditCardRegex = regexp.MustCompile(`\d{4}[ -]?\d{4}[ -]?\d{4}[ -]?\d{4}|\d{4}[ -]?\d{6}[ -]?\d{4}\d?`)
}

// redactResult is the result of a redaction process for any type of PII.
// Includes data about the redaction that occurred which can be useful for
// metrics.
type redactResult struct {
	// The number of times a substring in the source string was redacted.
	numRedacted int64
	// The string with all target substrings removed by the redaction process.
	redacted []byte
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

func main() {
	creditCardNumber := "4242424242424242"
	sourceString := fmt.Sprintf("My credit card number is %s. Please don't log it!", creditCardNumber)

	res := redactCreditCard([]byte(sourceString))

	if res.numRedacted == 0 {
		fmt.Printf("Nothing to redact. Source string:\n%s\n", sourceString)
	} else {
		fmt.Printf("Redacted credit card number %d time(s). Processed string:\n%s\n", res.numRedacted, res.redacted)
	}
}
