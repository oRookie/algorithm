package rumen

import "testing"

/**
最大子序和
输入: [-2,1,-3,4,-1,2,1,-5,4],
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6
int sum=nums[0];
        int n=nums[0];
        for(int i=1;i<nums.length;i++) {
            if(n>0)
				n+=nums[i];
            else
				n=nums[i];
            if(sum<n)
 				sum=n;
        }
        return sum;
*/

func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		nums[i] = dp(nums[i], nums[i-1])
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

func dp(a, b int) int {
	if a+b > a {
		return a + b
	}
	return a
}

func TestMaxSubArray(t *testing.T) {
	println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}


