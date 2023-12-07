package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

//You are given a 0-indexed integer array nums. Rearrange the values of nums
//according to the following rules:
//
//
// Sort the values at odd indices of nums in non-increasing order.
//
//
//
// For example, if nums = [4,1,2,3] before this step, it becomes [4,3,2,1]
//after. The values at odd indices 1 and 3 are sorted in non-increasing order.
//
//
// Sort the values at even indices of nums in non-decreasing order.
//
// For example, if nums = [4,1,2,3] before this step, it becomes [2,1,4,3]
//after. The values at even indices 0 and 2 are sorted in non-decreasing order.
//
//
//
//
// Return the array formed after rearranging the values of nums.
//
//
// Example 1:
//
//
//Input: nums = [4,1,2,3]
//Output: [2,3,4,1]
//Explanation:
//First, we sort the values present at odd indices (1 and 3) in non-increasing
//order.
//So, nums changes from [4,1,2,3] to [4,3,2,1].
//Next, we sort the values present at even indices (0 and 2) in non-decreasing
//order.
//So, nums changes from [4,1,2,3] to [2,3,4,1].
//Thus, the array formed after rearranging the values is [2,3,4,1].
//
//
// Example 2:
//
//
//Input: nums = [2,1]
//Output: [2,1]
//Explanation:
//Since there is exactly one odd index and one even index, no rearrangement of
//values takes place.
//The resultant array formed is [2,1], which is the same as the initial array.
//
//
//
// Constraints:
//
//
// 1 <= nums.length <= 100
// 1 <= nums[i] <= 100
//
//
// Related Topics Array Sorting 👍 685 👎 55

// leetcode submit region begin(Prohibit modification and deletion)
func sortEvenOdd(nums []int) []int {
	var oddList []int
	var evenList []int
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			evenList = append(evenList, nums[i])
		} else {
			oddList = append(oddList, nums[i])
		}
	}

	//sort.Ints(oddList)
	sort.Slice(oddList, func(i, j int) bool {
		return oddList[i] > oddList[j]
	})
	sort.Ints(evenList)
	ret := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			ret[i] = evenList[i/2]
		} else {
			ret[i] = oddList[i/2]
		}
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)

func Test_SortEvenAndOddIndicesIndependently(t *testing.T) {
	ret := sortEvenOdd([]int{4, 1, 2, 3})
	fmt.Println(ret)
}

func Test_SortEvenAndOddIndicesIndependently_table(t *testing.T) {
	type args struct {
	}

	tests := []struct {
		name     string
		args     args
		expected interface{}
	}{
		{},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

		})
	}

}
