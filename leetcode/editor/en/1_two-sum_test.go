package leetcode

import (
	"testing"
)

//Given an array of integers nums and an integer target, return indices of the
//two numbers such that they add up to target.
//
// You may assume that each input would have exactly one solution, and you may
//not use the same element twice.
//
// You can return the answer in any order.
//
//
// Example 1:
//
//
//Input: nums = [2,7,11,15], target = 9
//Output: [0,1]
//Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
//
//
// Example 2:
//
//
//Input: nums = [3,2,4], target = 6
//Output: [1,2]
//
//
// Example 3:
//
//
//Input: nums = [3,3], target = 6
//Output: [0,1]
//
//
//
// Constraints:
//
//
// 2 <= nums.length <= 10⁴
// -10⁹ <= nums[i] <= 10⁹
// -10⁹ <= target <= 10⁹
// Only one valid answer exists.
//
//
//
//Follow-up: Can you come up with an algorithm that is less than
//O(n²)
// time complexity?
//
// Related Topics Array Hash Table 👍 53436 👎 1770

// leetcode submit region begin(Prohibit modification and deletion)
func twoSum(nums []int, target int) []int {
	//make it map for sorting
	numsMap := make(map[int]int)
	for i, v := range nums {
		numsMap[v] = i
	}

	for i, v := range nums {

		t := target - v
		if index, exists := numsMap[t]; exists {
			if index == i {
				continue
			}
			return []int{i, index}
		}
	}
	return nil

}

//leetcode submit region end(Prohibit modification and deletion)

func Test_TwoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}

	tests := []struct {
		name     string
		args     args
		expected []int
	}{
		{
			name: "1",
			args: args{
				nums:   []int{2, 7, 11, 5},
				target: 9,
			},
			expected: []int{0, 1},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Log(twoSum(test.args.nums, test.args.target))
		})
	}

}
