package helper

import (
	b64 "encoding/base64"
)

/**
 *
 */
func Decode(encoded string) string {
	byteArray, _ := b64.StdEncoding.DecodeString(encoded)
	return string(byteArray[:])
}

/**
 *
 */
func Encode(decoded string) string {
	return b64.StdEncoding.EncodeToString([]byte(decoded))
}
