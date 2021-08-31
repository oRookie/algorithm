package rumen

import (
	"strings"
	"testing"
)

/**
输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
*/
func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	n := len(s)
	rk, ans := 0, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			// 左指针向右移动一格，移除一个字符
			delete(m, s[i-1])
		}
		for rk < n && m[s[rk]] == 0 {
			// 不断地移动右指针
			m[s[rk]]++
			rk++
		}
		// 第 i 到 rk 个字符是一个极长的无重复字符子串
		ans = max(ans, rk-i)
	}
	return ans
}

func lengthOfLongestSubstring1(s string) int {
	m := make(map[byte]int)
	n := len(s)
	j, ans := 0, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			delete(m, s[i-1])
		}
		for ; j < n && m[s[j]] == 0; j++ {
			m[s[j]]++
		}
		ans = max(ans, j-i)
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func TestLlengthOfLongestSubstring(t *testing.T) {
	println(lengthOfLongestSubstring1("pwwekew"))
	println(lengthOfLongestSubstring1("abcabcbb"))
}

/**
输入：s1 = "ab" s2 = "eidbaooo"
输出：true
解释：s2 包含 s1 的排列之一 ("ba").
*/
func checkInclusion(s1 string, s2 string) bool {
	n := len(s2) - 1
	for i := 0; i < n; i++ {
		j := i + len(s1)
		if strings.Compare(s1, s2[i:j]) == 0 {
			return true
		}
	}
	return false
}

func TestCcheckInclusion(t *testing.T) {
	println(checkInclusion("ba", "eidbaooo"))
}
