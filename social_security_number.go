package main

import (
	"regexp"
)

const (
	ninUSRedactStub = "PII_NIN_US"
)

var ninUSRegex *regexp.Regexp

func init() {
	// Regex source:
	// https://github.com/solvvy/redact-pii/blob/main/src/built-ins/simple-regexp-patterns.ts

	// Using terminology "national identification number", by country.
	// https://en.wikipedia.org/wiki/National_identification_number
	ninUSRegex = regexp.MustCompile(`\b\d{3}[ -.]\d{2}[ -.]\d{4}\b`)
}

// redactNinUS redacts national identification numbers for the US
// (social security numbers/SSNs) numbers from strings.
func redactNinUS(src []byte) redactResult {
	numRedacted := int64(0)

	processed := ninUSRegex.ReplaceAllFunc(src, func(_ []byte) []byte {
		numRedacted += 1
		return []byte(ninUSRedactStub)
	})

	return redactResult{
		numRedacted: numRedacted,
		redacted:    processed,
	}
}
