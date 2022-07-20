package confx

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"
)

func ResolveProperties(path string) (map[string]string, error) {
	file, _ := os.Open(path)
	defer file.Close()
	_, err := file.Stat()
	if err != nil {
		return nil, errors.New("file is not exist")
	}
	conf := make(map[string]string)
	br := bufio.NewReader(file)
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			// 读取结束
			break
		}
		if err != nil {
			return nil, err
		}
		lineContent := string(line)
		prop := strings.TrimSpace(lineContent)
		if prop == "" {
			continue
		}
		key := prop[:strings.Index(prop, "=")]
		val := prop[strings.Index(prop, "=")+1:]
		conf[key] = val
	}
	return conf, nil
}
