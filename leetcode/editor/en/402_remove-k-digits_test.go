package leetcode

import (
	"strings"
	"testing"
)

//Given string num representing a non-negative integer num, and an integer k,
//return the smallest possible integer after removing k digits from num.
//
//
// Example 1:
//
//
//Input: num = "1432219", k = 3
//Output: "1219"
//Explanation: Remove the three digits 4, 3, and 2 to form the new number 1219
//which is the smallest.
//
//
// Example 2:
//
//
//Input: num = "10200", k = 1
//Output: "200"
//Explanation: Remove the leading 1 and the number is 200. Note that the output
//must not contain leading zeroes.
//
//
// Example 3:
//
//
//Input: num = "10", k = 2
//Output: "0"
//Explanation: Remove all the digits from the number and it is left with
//nothing which is 0.
//
//
//
// Constraints:
//
//
// 1 <= k <= num.length <= 10⁵
// num consists of only digits.
// num does not have any leading zeros except for the zero itself.
//
//
// Related Topics String Stack Greedy Monotonic Stack 👍 8349 👎 364

// leetcode submit region begin(Prohibit modification and deletion)
func removeKdigits(num string, k int) string {

	r := []rune(num)
	for {
		if k == 0 {
			// Remove head 0s
			num = strings.TrimLeft(string(r), "0")
			if num == "" {
				return "0"
			}
			return num
		}

		if len(r) == 1 {
			return "0"
		}

		removeOne := false
		for index, _ := range r {
			if index == 0 {
				continue
			}
			if r[index-1] > r[index] {
				r = append(r[0:index-1], r[index:]...)
				removeOne = true
				break
			}
		}
		if removeOne == false {
			r = append(r[0 : len(r)-1])
		}
		k -= 1
	}

}

//leetcode submit region end(Prohibit modification and deletion)

func Test_RemoveKDigits(t *testing.T) {
	type args struct {
		arg1 string
		arg2 int
	}

	tests := []struct {
		name     string
		args     args
		expected string
	}{
		{
			name:     "1",
			args:     args{arg1: "1432219", arg2: 3},
			expected: "1219",
		},
		{
			name:     "2",
			args:     args{arg1: "10200", arg2: 1},
			expected: "200",
		},
		{
			name:     "3",
			args:     args{arg1: "10", arg2: 2},
			expected: "0",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if result := removeKdigits(test.args.arg1, test.args.arg2); result != test.expected {
				t.Errorf("not match, expect : %v, actual : %v", test.expected, result)
			}
		})
	}
}
