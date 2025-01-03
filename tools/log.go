package tools

import (
	"fmt"
	"github.com/fatih/color"
	"log"
)

type Log struct {
}

func (*Log) ERR(err error) {
	color.Red("Error: %s", err.Error())
}

func (*Log) INFO(s string) {
	log.Println(fmt.Sprintf("Info: %s", s))
}

func (*Log) WARN(s string) {
	color.Yellow("Warn: %s", s)
}

func (*Log) MARK(s string) {
	color.Green("%s", s)
}
