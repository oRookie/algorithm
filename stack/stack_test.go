package stack

import (
	"strconv"
	"strings"
	"testing"
)

func TestArrayStack(t *testing.T) {
	stack := NewArrayStack(10)
	for i := 1; i <= 5; i++ {
		stack.Push(strconv.Itoa(i))
	}
	println(stack.Pop())
	println(stack.Pop())
	println(stack.Pop())
}

func TestCharMatching(t *testing.T) {
	println(checkValid("[{()()}]"))
}

/**
假设表达式中只包含三种括号，圆括号 ()、方括号 [] 和花括号{}，并且它们可以任意嵌套。比如，{[{}]}或 [{()}([])] 等都为合法格式，而{[}()] 或 [({)] 为不合法的格式。对于一个包含三种括号的表达式字符串，如何检查它是否合法呢？
*/
var s1, s2, s3, sl, sr = "}", ")", "]", "{([", "})]"

func checkValid(str string) bool {
	stack := NewArrayStack(len(str) / 2 + 1)
	for _, v := range str {
		if strings.Contains(sl, string(v)) {
			stack.Push(string(v))
		}
		if strings.Contains(sr, string(v)) {
			if compare(stack.Pop(),string(v)){
				continue
			}else{
				return false
			}
		}
	}
	return true
}

func compare(sl, sr string) bool {
	switch sr {
	case s1:
		return strings.Compare(sl, "{") == 0
	case s2:
		return strings.Compare(sl, "(") == 0
	case s3:
		return strings.Compare(sl, "[") == 0
	}
	return false
}

type arrayStack struct {
	items []string
	cap   int
	len   int
}

func NewArrayStack(cap int) *arrayStack {
	var item []string
	return &arrayStack{
		items: item,
		cap:   cap,
		len:   0,
	}
}

func (s *arrayStack) Push(a string) bool {
	if s.len >= s.cap {
		return false
	}
	s.items = append(s.items, a)
	s.len++
	return true
}

func (s *arrayStack) Pop() string {
	if s.len == 0 {
		return ""
	}
	tmp := s.items[s.len-1]
	s.len--
	return tmp
}
