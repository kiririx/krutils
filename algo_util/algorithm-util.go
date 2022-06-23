package algo_util

import (
	"crypto/md5"
	"encoding/base64"
	"errors"
	"fmt"
)

func Base64Encode(v string) string {
	return base64.StdEncoding.EncodeToString([]byte(v))
}

func Base64Decode(v string) (string, error) {
	b, err := base64.StdEncoding.DecodeString(v)
	if err != nil {
		return "", errors.New("decoding fail")
	}
	return string(b), nil
}

func MD5(v string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(v)))
}
