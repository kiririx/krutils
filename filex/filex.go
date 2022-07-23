package filex

import (
	"errors"
	"path"
	"regexp"
	"strings"
)

// GetUrlFileExt returns the file extension of the url.
//
// for example, input (http://xx?a=1&file=yourfile.png&b=2, ['png','jpg']), the function will return png
func GetUrlFileExt(filename string, allowExt []string) (string, error) {
	fileName := path.Base(filename)
	reg, err := regexp.Compile("\\.(" + strings.Join(allowExt, "|") + ")")
	if err != nil {
		return "", err
	}
	matchedExtArr := reg.FindAllString(fileName, -1)
	if matchedExtArr != nil && len(matchedExtArr) > 0 {
		ext := matchedExtArr[len(matchedExtArr)-1]
		return ext[1:], nil
	}
	return "", errors.New("获取文件扩展名失败")
}
