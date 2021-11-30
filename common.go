package main

// redactResult is the result of a redaction process for any type of PII.
// Includes data about the redaction that occurred which can be useful for
// metrics.
type redactResult struct {
	// The number of times a substring in the source string was redacted.
	numRedacted int64
	// The string with all target substrings removed by the redaction process.
	redacted []byte
}
