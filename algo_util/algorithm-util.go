package algo_util

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
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

func Sha256(v string) string {
	h := sha256.New()
	h.Write([]byte(v))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func UUID() string {
	v := uuid.NewV4()
	return fmt.Sprintf("%s", v)
}
