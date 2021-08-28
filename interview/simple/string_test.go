package test

import (
	"strings"
	"testing"
)

/**
字符串匹配
在utf8字符串判断是否包含指定字符串，并返回下标。 “北京天安门最美丽” , “天安门” 结果：2
*/
func TestFindIndex(t *testing.T) {
	println(findIndex("北京天安门最美丽", "天安门"))
}

func findIndex(str string, substr string) int {
	asciiPos := strings.Index(str, substr)
	if asciiPos == -1 || asciiPos == 0 {
		return asciiPos
	}
	pos := 0
	totalSize := 0
	reader := strings.NewReader(str)
	for _, size, err := reader.ReadRune(); err == nil; _, size, err = reader.ReadRune() {
		totalSize += size
		pos++
		// 匹配到
		if totalSize == asciiPos {
			return pos
		}
	}
	return pos
}
