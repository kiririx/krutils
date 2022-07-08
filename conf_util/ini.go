package conf_util

import (
	"github.com/kiririx/krutils/str_util"
	"io"
	"os"
	"regexp"
	"strings"
)

func ResolveIni(file *os.File) (Ini, error) {
	NilIni := Ini{}
	_, err := file.Stat()
	if err != nil {
		return NilIni, err
	}
	var buf [128]byte
	sections := make(map[string]*Properties)
	currentSection := ""
	for {
		n, err := file.Read(buf[:])
		if err == io.EOF {
			// 读取结束
			break
		}
		if err != nil {
			return NilIni, err
		}
		lineContent := string(buf[:n])
		// todo: struct复制
		lineContent = str_util.TrimSpace(lineContent)
		if lineContent != "" {
			if ok, _ := regexp.MatchString("(\\[).*(\\])", lineContent); ok {
				// sections
				lineContent = strings.TrimPrefix(lineContent, "[")
				lineContent = strings.TrimSuffix(lineContent, "]")
				sections[lineContent] = &Properties{}
				currentSection = lineContent
			} else if ok, _ := regexp.MatchString(".*=.*", lineContent); ok {
				// k v
				contentArr := strings.Split(lineContent, "=")
				sections[currentSection].Set(contentArr[0], contentArr[1])
			}
		}
	}
	iniFile := Ini{Sections: sections}
	return iniFile, nil
}
