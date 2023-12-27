package converter

import "encoding/base64"

func ToBase64(file []byte) string {
	return base64.StdEncoding.EncodeToString(file)
}
