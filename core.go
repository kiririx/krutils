package krutils

import (
	"crypto/md5"
	"encoding/base64"
	"errors"
	"fmt"
)

func ToInt(v string) int {
	return 0
}

func encodeBase64(v string) string {
	return base64.StdEncoding.EncodeToString([]byte(v))
}

func decodeBase64(v string) (string, error) {
	b, err := base64.StdEncoding.DecodeString(v)
	if err != nil {
		return "", errors.New("decoding fail")
	}
	return string(b), nil
}

func MD5(v string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(v)))
}
