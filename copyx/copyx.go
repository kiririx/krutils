package copyx

import "github.com/kiririx/krutils/structx"

func DeepCopyStruct(src any, target any) error {
	return structx.DeepCopy(src, target)
}
