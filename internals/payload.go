package internals

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"

	"github.com/PS-Wizard/JWT/internals/encryptions"
)

func generateRandomID() string {
	bytes := make([]byte, 8)
	_, err := rand.Read(bytes[:])
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(bytes)
}

func (j *JWTConfig) GeneratePayload() string {
	j.Claims.Id = generateRandomID()
	payload, err := json.Marshal(j.Claims)
	if err != nil {
		panic(err)
	}
	return encryptions.EncodeBase64(string(payload))
}
