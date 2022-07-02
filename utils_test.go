package krutils

import (
	"github.com/kiririx/krutils/str_util"
	"testing"
)

func TestConv(t *testing.T) {
	t.Log(str_util.Contains("张三", "三", "c"))
}

type ST struct {
	a int
}
