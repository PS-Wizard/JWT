package encryptions

import "encoding/base64"

func EncodeBase64(x string) string {
	return base64.StdEncoding.EncodeToString([]byte(x))
}

