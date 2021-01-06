package strlib

import (
	"strings"
	"unicode/utf8"
)

//Utf8Index 替换 strings.Index 处理 Unicode 值
func Utf8Index(str, substr string) int {
	index := strings.Index(str, substr)
	if index < 0 {
		return -1
	}
	return utf8.RuneCountInString(str[:index])
}
