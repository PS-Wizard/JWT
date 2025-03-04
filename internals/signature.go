package internals

import (
	"crypto/hmac"
	"crypto/sha256"

	"github.com/PS-Wizard/JWT/internals/encryptions"
)

func (j *JWTConfig) GenerateSignature() string {
	header := j.GenerateHeader()
	payload := j.GeneratePayload()

	h := hmac.New(sha256.New, []byte(j.Secret))

	h.Write([]byte(header + "." + payload))
	signature := h.Sum(nil)
	return encryptions.EncodeBase64(string(signature))
}
