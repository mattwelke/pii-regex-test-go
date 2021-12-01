package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_redactCreditCard(t *testing.T) {
	// Fixture source: https://stripe.com/docs/testing
	testCreditCardNumbers := []string{
		// Visa
		string("4242424242424242"),
		string("4000056655665556"),

		// Mastercard
		string("5555555555554444"),
		string("2223003122003222"),
		string("5200828282828210"),
		string("5105105105105100"),

		// American Express
		string("378282246310005"),
		string("371449635398431"),

		// Discover
		string("6011111111111117"),
		string("6011000990139424"),

		// Diners Club
		string("3056930009020004"),
		string("36227206271667"),

		// JCB
		string("3566002020360505"),

		// UnionPay
		string("6200000000000005"),
	}

	testCreditCardNumbersStylized := []string{
		// Visa
		string("4242 4242 4242 4242"),
		string("4000 0566 5566 5556"),

		// Mastercard
		string("5555 5555 5555 4444"),
		string("2223 0031 2200 3222"),
		string("5200 8282 8282 8210"),
		string("5105 1051 0510 5100"),

		// American Express
		string("3782 822463 10005"),
		string("3782 822463 10005"),

		// Discover
		string("6011 1111 1111 1117"),
		string("6011 0009 9013 9424"),

		// Diners Club
		string("3056 9300 0902 0004"),
		string("3622 720627 1667"),

		// JCB
		string("3566 0020 2036 0505"),

		// UnionPay
		string("6200 0000 0000 0005"),
	}

	allCreditCardNumbers := removeDuplicateStr(append(testCreditCardNumbers, testCreditCardNumbersStylized...))

	for _, ccNum := range allCreditCardNumbers {
		type args struct {
			src []byte
		}
		tests := []struct {
			name string
			args args
			want redactResult
		}{

			{
				name: "Can sanitize a string that is a credit card number, reporting 1 redaction.",
				args: args{[]byte(ccNum)},
				want: redactResult{
					numRedacted: 1,
					redacted:    []byte(creditCardRedactStub),
				},
			},
			{
				name: "Can sanitize a string contains a credit card number, reporting 1 redaction.",
				args: args{[]byte(fmt.Sprintf(" %s ", ccNum))},
				want: redactResult{
					numRedacted: 1,
					redacted:    []byte(fmt.Sprintf(" %s ", creditCardRedactStub)),
				},
			},
			{
				name: "Can sanitize a string that contains multiple credit card numbers, reporting 2 redactions.",
				args: args{[]byte(fmt.Sprintf("%s %s", ccNum, ccNum))},
				want: redactResult{
					numRedacted: 2,
					redacted:    []byte(fmt.Sprintf("%s %s", creditCardRedactStub, creditCardRedactStub)),
				},
			},
			{
				name: "Can sanitize a string that contains multiple credit card numbers, separated by a newline, reporting 2 redactions.",
				args: args{[]byte(fmt.Sprintf("%s\n%s", ccNum, ccNum))},
				want: redactResult{
					numRedacted: 2,
					redacted:    []byte(fmt.Sprintf("%s\n%s", creditCardRedactStub, creditCardRedactStub)),
				},
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := redactCreditCard(tt.args.src); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("redactCreditCard() = %+v, want %+v",
						redactResultForDisplayFromRedactResult(got),
						redactResultForDisplayFromRedactResult(tt.want))
				}
			})
		}
	}
}
