package logx

import (
	"fmt"
	"github.com/fatih/color"
	"log"
)

func ERR(err error) {
	color.Red("Error: %s", err.Error())
}

func INFO(s string) {
	log.Println(fmt.Sprintf("Info: %s", s))
}

func WARN(s string) {
	color.Yellow("Warn: %s", s)
}

func MARK(s string) {
	color.Green("%s", s)
}
