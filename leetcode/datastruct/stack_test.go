package datastruct

import (
	"fmt"
	"sort"
	"testing"
)

func isValid(s string) bool {
	m := map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
	}
	c := []byte(s)
	hp := &hp{S: []string{}}
	for _, v := range c {
		if len(s)%2 == 1 {
			return false
		}
		v := string(v)
		if v == "(" || v == "{" || v == "[" {
			hp.Push(v)
		} else {
			vv := hp.Pop()
			if v != m[vv] {
				return false
			}
		}
	}
	if len(hp.S) == 0 {
		return true
	}
	return false
}

type hp struct {
	S []string
}

func (h *hp) Push(v string) {
	h.S = append(h.S, v)
}

func (h *hp) Pop() string {
	a := h.S
	if len(a) == 0 {
		return ""
	}
	v := a[len(a)-1]
	h.S = a[:len(a)-1]
	return v
}

func TestMerge(t *testing.T) {
	num1 := []int{1, 2, 3, 0, 0, 0}
	merge(num1, 3, []int{2, 5, 6, 0, 0}, 3)
	fmt.Printf("%v", num1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	num := sort.IntSlice{}
	num = append(nums1[:m], nums2[:n]...)
	num.Sort()
	nums1 = num
}
