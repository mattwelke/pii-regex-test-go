package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_redactNinUS(t *testing.T) {
	// There is no list of "test SSNs" but according to the rules at
	// https://secure.ssa.gov/poms.nsf/lnx/0110201035, we can infer some SSNs
	// that some people might have. So these are possibly in use SSNs here, but
	// they are not considered PII because they aren't linked to any data we
	// collected. They're just code for our tests here.
	testSocialSecurityNumbers := removeDuplicateStr([]string{
		"001-01-0001",
	})

	for _, ssn := range testSocialSecurityNumbers {
		type args struct {
			src []byte
		}
		tests := []struct {
			name string
			args args
			want redactResult
		}{

			{
				name: "Can sanitize a string that is a social security number, reporting 1 redaction.",
				args: args{[]byte(ssn)},
				want: redactResult{
					numRedacted: 1,
					redacted:    []byte(ninUSRedactStub),
				},
			},
			{
				name: "Can sanitize a string contains a social security number, reporting 1 redaction.",
				args: args{[]byte(fmt.Sprintf(" %s ", ssn))},
				want: redactResult{
					numRedacted: 1,
					redacted:    []byte(fmt.Sprintf(" %s ", ninUSRedactStub)),
				},
			},
			{
				name: "Can sanitize a string that contains multiple social security numbers, reporting 2 redactions.",
				args: args{[]byte(fmt.Sprintf("%s %s", ssn, ssn))},
				want: redactResult{
					numRedacted: 2,
					redacted:    []byte(fmt.Sprintf("%s %s", ninUSRedactStub, ninUSRedactStub)),
				},
			},
			{
				name: "Can sanitize a string that contains multiple social security numbers, separated by a newline, reporting 2 redactions.",
				args: args{[]byte(fmt.Sprintf("%s\n%s", ssn, ssn))},
				want: redactResult{
					numRedacted: 2,
					redacted:    []byte(fmt.Sprintf("%s\n%s", ninUSRedactStub, ninUSRedactStub)),
				},
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := redactNinUS(tt.args.src); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("redactNinUS() = %+v, want %+v",
						redactResultForDisplayFromRedactResult(got),
						redactResultForDisplayFromRedactResult(tt.want))
				}
			})
		}
	}
}
