package logx

import (
	"fmt"
	"log"
)

func ERR(err error) {
	log.Println(fmt.Sprintf(`Err: %s`, err.Error()))
}

func INFO(s string) {
	log.Println(fmt.Sprintf("Info: %s", s))
}

func WARN(s string) {
	log.Println(fmt.Sprintf("Warn: %s", s))
}
