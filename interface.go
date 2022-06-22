package krutils

type base64Util struct{}
type strUtil struct{}
type md5Util struct{}

var Base64Util base64Util
var StrUtil strUtil
var Md5Util md5Util

func init() {

}

func (*base64Util) Encode(v string) string {
	return encodeBase64(v)
}

func (*base64Util) Decode(v string) (string, error) {
	return decodeBase64(v)
}

func (*strUtil) IntToStr(v int64) string {
	return NumToStr(v)
}

func (*strUtil) FloatToStr(v float64) string {
	return NumToStr(v)
}

func (*strUtil) ToStr(v any) string {
	return toStr(v)
}

func (*md5Util) Md5(v string) string {
	return MD5(v)
}
