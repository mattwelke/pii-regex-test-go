package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_redactPhoneNumber(t *testing.T) {
	testPhoneNumbers := removeDuplicateStr([]string{
		// Source: https://support.twilio.com/hc/en-us/articles/223183008-Formatting-International-Phone-Numbers
		"(415) 555-2671",

		// Variations:
		"415-555-2671",
		"415.555.2671",
		"415 555 2671",
		"1 415 555 2671",
		"+14155552671",
	})

	for _, phoneNumStr := range testPhoneNumbers {
		phoneNum := []byte(phoneNumStr)

		type args struct {
			src []byte
		}
		tests := []struct {
			name string
			args args
			want redactResult
		}{

			{
				name: "Can sanitize a string that is a phone number, reporting 1 redaction.",
				args: args{phoneNum},
				want: redactResult{
					numRedacted: 1,
					redacted:    []byte(phoneNumberRedactStub),
				},
			},
			{
				name: "Can sanitize a string contains a phone number, reporting 1 redaction.",
				args: args{[]byte(fmt.Sprintf(" %s ", phoneNum))},
				want: redactResult{
					numRedacted: 1,
					redacted:    []byte(fmt.Sprintf(" %s ", phoneNumberRedactStub)),
				},
			},
			{
				name: "Can sanitize a string that contains multiple phone numbers, reporting 2 redactions.",
				args: args{[]byte(fmt.Sprintf("%s %s", phoneNum, phoneNum))},
				want: redactResult{
					numRedacted: 2,
					redacted:    []byte(fmt.Sprintf("%s %s", phoneNumberRedactStub, phoneNumberRedactStub)),
				},
			},
			{
				name: "Can sanitize a string that contains multiple phone numbers, separated by a newline, reporting 2 redactions.",
				args: args{[]byte(fmt.Sprintf("%s\n%s", phoneNum, phoneNum))},
				want: redactResult{
					numRedacted: 2,
					redacted:    []byte(fmt.Sprintf("%s\n%s", phoneNumberRedactStub, phoneNumberRedactStub)),
				},
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := redactPhoneNumber(tt.args.src); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("redactPhoneNumber() = %+v, want %+v",
						redactResultForDisplayFromRedactResult(got),
						redactResultForDisplayFromRedactResult(tt.want))
				}
			})
		}
	}
}

func Test_redactPhoneNumberNoProductIDFalsePositives(t *testing.T) {
	for _, knownProductIDStr := range knownProductIDs {
		knownProductID := []byte(knownProductIDStr)

		type args struct {
			src []byte
		}
		tests := []struct {
			name string
			args args
			want redactResult
		}{

			{
				name: "Can sanitize a string that is a known product ID from one of our customers without falsely considering it to be a phone number, reporting no redactions.",
				args: args{knownProductID},
				want: redactResult{
					numRedacted: 0,
					redacted:    []byte(knownProductID),
				},
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := redactPhoneNumber(tt.args.src); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("redactPhoneNumber() = %+v, want %+v",
						redactResultForDisplayFromRedactResult(got),
						redactResultForDisplayFromRedactResult(tt.want))
				}
			})
		}
	}
}
