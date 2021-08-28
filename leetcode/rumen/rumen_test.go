package rumen

import "testing"

/**
给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。


示例 1:

输入: nums = [-1,0,3,5,9,12], target = 9
输出: 4
解释: 9 出现在 nums 中并且下标为 4
示例 2:

输入: nums = [-1,0,3,5,9,12], target = 2
输出: -1
解释: 2 不存在 nums 中因此返回 -1

*/
func search(nums []int, target int) int {
	index := -1
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target { // you
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			index = mid
			return index
		}
	}
	return index
}

/**
输入：n = 5, bad = 4
输出：4
解释：
调用 isBadVersion(3) -> false
调用 isBadVersion(5) -> true
调用 isBadVersion(4) -> true
所以，4 是第一个错误的版本。

找第一个错误版本
*/
func findFirstErrorVersion(n, bad int) int {
	ans := 1
	left, right := 1, n
	for left < right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid
			ans = mid
		} else {
			left = left + 1
		}
	}
	return ans
	//left,right := 1,n
	//for left < right {
	//	mid := left + (right - left) / 2
	//	if isBadVersion(mid) { //答案在left-mid中
	//		right = mid
	//	}else { //答案在mid+1 - right中
	//		left = mid + 1
	//	}
	//}
	//return left
}

func isBadVersion(i int) bool {
	if i == 2 {
		return true
	}
	return false
}

/**
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
求大于等于target的index值
请必须使用时间复杂度为 O(log n) 的算法。
输入: nums = [1,3,5,6], target = 2
输出: 1
ans := len(nums)
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if target <= nums[mid] {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
*/

func searchInsert(nums []int, target int) int {
	ans := len(nums)
	left, right := 0, ans-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] >= target {
			right = mid - 1
			ans = mid
		} else {
			left = mid + 1
		}
	}
	return ans
}

func TestBinarySearch(t *testing.T) {
	nums := []int{1,3,5,6}
	println(searchInsert(nums, 3))

}
