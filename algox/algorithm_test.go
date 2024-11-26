package algox

import "testing"

func TestBase64Encode(t *testing.T) {
	t.Log(Base64Encode("1234567"))
}

func TestBase64Decode(t *testing.T) {
	t.Log(Base64Decode("MTIzNDU2Nw=="))
}

func TestMD5(t *testing.T) {
	t.Log(MD5("1"))
}

func TestSHA256(t *testing.T) {
	t.Log(Sha256("1"))
}

func TestRandomInt(t *testing.T) {
	for range 10 {
		t.Log(RandomInt(0, 101111111))
	}
}

func TestUUID2(t *testing.T) {
	t.Log(UUID())
}
