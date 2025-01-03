package tools

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"time"
)

type Algorithm struct {
}

func (receive *Algorithm) Base64Encode(v string) string {
	return base64.StdEncoding.EncodeToString([]byte(v))
}

func (receive *Algorithm) Base64Decode(v string) (string, error) {
	b, err := base64.StdEncoding.DecodeString(v)
	if err != nil {
		return "", errors.New("decoding fail")
	}
	return string(b), nil
}

func (receive *Algorithm) MD5(v string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(v)))
}

func (receive *Algorithm) Sha256(v string) string {
	h := sha256.New()
	h.Write([]byte(v))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func (receive *Algorithm) UUID() string {
	v := uuid.NewV4()
	return fmt.Sprintf("%s", v)
}

func (receive *Algorithm) RandomInt(start, end int) int {
	end += 1
	if start < 0 || end < 0 {
		return 0
	}
	if start == end {
		return start
	}
	if start > end {
		return end
	}
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(end-start) + start
}
