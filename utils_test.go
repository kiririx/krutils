package krutils

import (
	"github.com/kiririx/krutils/str_util"
	"testing"
)

func TestConv(t *testing.T) {
	t.Log(str_util.ToStr(make(chan int)))
}
