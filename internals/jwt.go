package internals

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"
)

const secret string = "69420lmao"

func Base64_Encode(someData string) string {
	encoded := base64.URLEncoding.EncodeToString([]byte(someData))
	return encoded
}

func GenerateHeader() string {
	header := map[string]any{
		"alg": "HS256",
		"typ": "JWT",
	}

	headerJSON, err := json.Marshal(header)
	if err != nil {
		fmt.Println("Error marshalling header:", err)
		return ""
	}

	return Base64_Encode(string(headerJSON))
}

func GenerateSignature(encodedHeader, encodedPayload string) string {
	h := hmac.New(sha256.New, []byte(secret))
	data := encodedHeader + "." + encodedPayload
	h.Write([]byte(data))
	return base64.URLEncoding.EncodeToString(h.Sum(nil))
}

func (user *User) GenerateJWT() {
	payload := map[string]any{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 1).Unix(), // 1 hour expiration
	}

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshalling payload:", err)
		return
	}
	encodedPayload := Base64_Encode(string(payloadJSON))
	encodedHeader := GenerateHeader()
	signature := GenerateSignature(encodedHeader, encodedPayload)

	// Print the final JWT
	fmt.Println("JWT: ", fmt.Sprintf("%s.%s.%s", encodedHeader, encodedPayload, signature))
}
