package tests

import (
	"testing"

	"github.com/PS-Wizard/JWT/internals/encryptions"
)

type Base64Tests struct {
	input string
	want  string
}

var base64Tests = []Base64Tests{
	{"hello", "aGVsbG8="},
	{"world", "d29ybGQ="},
	{"Go is cool!", "R28gaXMgY29vbCE="},
	{"1234567890", "MTIzNDU2Nzg5MA=="},
	{"!@#$%^&*()", "IUAjJCVeJiooKQ=="},
	{"", ""},
}

func TestEncodeBase64(t *testing.T) {

	for i, v := range base64Tests {
		got := encryptions.EncodeBase64(v.input)
		if got != v.want {
			t.Errorf("Test %d failed: --Wanted %s, got %s: --", i, v.want, got)
		}
	}
}
