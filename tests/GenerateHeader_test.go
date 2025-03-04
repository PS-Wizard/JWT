package tests

import (
	"testing"

	"github.com/PS-Wizard/JWT/internals"
)

type JWTHeaderTest struct {
	config internals.JWTConfig
	want   string
}

var tests = []JWTHeaderTest{
	{
		config: internals.JWTConfig{
			Alg: "HS256",
			Typ: "JWT",
		},
		want: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9",
	},
	{
		config: internals.JWTConfig{
			Alg: "RS256",
			Typ: "JWT",
		},
		want: "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9",
	},
}

func TestGenerateHeader(t *testing.T) {
	for i, test := range tests {
		got := test.config.GenerateHeader()
		if got != test.want {
			t.Errorf("Test %d failed: --  Wanted: %s, Got: %s ---", i, test.want, got)
		}
	}
}
