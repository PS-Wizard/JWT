package encryptions

import "encoding/base64"

func EncodeBase64(x string) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString([]byte(x))
}
