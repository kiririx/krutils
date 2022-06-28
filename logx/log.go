package logx

import (
	"fmt"
	"log"
)

func ErrorLog(function, s string) {
	log.Println(fmt.Sprintf(`%s Err: %s`, function, s))
}

func Info(s string) {
	log.Println(fmt.Sprintf("Info: %s", s))
}

func Warn(s string) {
	log.Println(fmt.Sprintf("Warn: %s", s))
}
