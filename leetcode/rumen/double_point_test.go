package rumen

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

/**
给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。
输入：nums = [-4,-1,0,3,10]
输出：[0,1,9,16,100]
解释：平方后，数组变为 [16,1,0,9,100]
排序后，数组变为 [0,1,9,16,100]

-

*/

func sortedSquares(nums []int) []int {
	ans := make([]int, len(nums))
	i, j := 0, len(nums)-1
	for pos := len(nums) - 1; pos >= 0; pos-- {
		if v1, v2 := nums[i]*nums[i], nums[j]*nums[j]; v1 > v2 {
			ans[pos] = v1
			i++
		} else {
			ans[pos] = v2
			j--
		}
	}
	return ans
}

/**
给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。

进阶：

尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
你可以使用空间复杂度为 O(1) 的 原地 算法解决这个问题吗？
 

示例 1:

输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
解释:
7 - 2  (6+3)%7
6 - 1  (5+3)%7
5 - 0   (4+3)%7
4 - 6    (3+3)%7
向右旋转 1 步: [7,1,2,3,4,5,6]
向右旋转 2 步: [6,7,1,2,3,4,5]
向右旋转 3 步: [5,6,7,1,2,3,4]
示例 2:

输入：nums = [-1,-100,3,99], k = 2
输出：[3,99,-1,-100]
解释:
向右旋转 1 步: [99,-1,-100,3]
向右旋转 2 步: [3,99,-1,-100]
 

提示：

1 <= nums.length <= 2 * 104
-231 <= nums[i] <= 231 - 1
0 <= k <= 105

*/
func rotate(nums []int, k int) {
	ans := make([]int, len(nums))
	for i, v := range nums {
		ans[(i+k)%len(nums)] = v
	}
	copy(nums, ans)
}

func rotate1(nums []int, k int) {
	k %= len(nums)
	//全部反转
	reverse(nums)
	//[0,k]
	reverse(nums[:k])
	//[k]
	reverse(nums[k:])
}

func reverse(nums []int) {
	for i, j := 0, len(nums); i < j/2; i++ {
		nums[i], nums[j-1-i] = nums[j-1-i], nums[i]
	}
}

func TestSortedSquares(t *testing.T) {
	nums := []int{-1, 1, 2}
	rotate1(nums, 4)
	t.Logf("%v", nums)
}

/**
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
输入: [0,1,0,3,12]
输出: [1,3,12,0,0]
*/
func moveZeroes(nums []int) {
	//ans := make([]int, len(nums))
	//n := len(nums) - 1
	//for i, j := 0, 0; i <= n; i++ {
	//	if nums[i] != 0 {
	//		ans[j] = nums[i]
	//		j++
	//	}
	//}
	//copy(nums, ans)
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

func TestMoveZeroes(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	t.Logf("%v", nums)
}

/**
给定一个已按照 升序排列  的整数数组 numbers ，请你从数组中找出两个数满足相加之和等于目标数 target 。

函数应该以长度为 2 的整数数组的形式返回这两个数的下标值。numbers 的下标 从 1 开始计数 ，所以答案数组应当满足 1 <= answer[0] < answer[1] <= numbers.length 。

你可以假设每个输入只对应唯一的答案，而且你不可以重复使用相同的元素。
输入：numbers = [2,7,11,15], target = 9
输出：[1,2]
解释：2 与 7 之和等于目标数 9 。因此 index1 = 1, index2 = 2

*/
func twoSum(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left]+nums[right] > target {
			right = right - 1
		} else if nums[left]+nums[right] < target {
			left = left + 1
		} else {
			return []int{left + 1, right + 1}
		}
	}
	return []int{}
}

func TestTwonums(t *testing.T) {
	fmt.Printf("%v", twoSum([]int{2}, 9))
}

func TestReserve(t *testing.T) {
	s := []byte("abcde")
	reverseString(s)
	println(string(s))
}

func Reserve(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

/**
输入：["h","e","l","l","o"]
输出：["o","l","l","e","h"]
*/
func reverseString(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

/**
输入："Let's take LeetCode contest"
输出："s'teL ekat edoCteeL tsetnoc"
*/
func reverseWords(s string) string {
	ret := []byte{}
	len := len(s)
	for i := 0; i < len; {
		start := i
		for i < len && s[i] != ' ' {
			i++
		}
		for p := start; p < i; p++ {
			ret = append(ret, s[start+i-1-p])
		}
		for ; i < len && s[i] == ' '; {
			i++
			ret = append(ret, ' ')
		}
	}
	return string(ret)
}

func TestReverseWords(t *testing.T) {
	println(reverseWords("Let's take LeetCode contes"))
}

/**
1,2,3,4,5,6,7,8,9  l
1,2,3,4,5,6,7,8,9,10,l+1

*/
func middleNode(head *ListNode) *ListNode {
	l, r := head, head
	for ; r.Next != nil && r.Next.Next != nil; {
		l, r = l.Next, r.Next.Next
	}
	if r.Next == nil {
		return l
	}
	if r.Next.Next == nil {
		return l.Next
	}
	return nil
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func TestMiddleNode(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val:  6,
							Next: nil,
						},
					},
				},
			},
		},
	}
	fmt.Printf("%+v", middleNode(head))
}

/**
 * 输入：head = [1,2,3,4,5], n = 2
 * 输出：[1,2,3,5]
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	i, j := head, head
	for i := 0; i < n; i++ {
		j = j.Next
	}
	if j == nil {
		head = head.Next
		return head
	}
	for j.Next != nil {
		i, j = i.Next, j.Next
	}
	i.Next = i.Next.Next
	return head
}

func TestRemoveNthFromEnd(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}
	LogJson(removeNthFromEnd(head, 2))
}

func LogJson(i interface{}) {
	bs, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	log.Println(string(bs))
}
