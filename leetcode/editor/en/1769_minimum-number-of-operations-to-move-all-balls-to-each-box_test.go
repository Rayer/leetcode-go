package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//You have n boxes. You are given a binary string boxes of length n, where
//boxes[i] is '0' if the iᵗʰ box is empty, and '1' if it contains one ball.
//
// In one operation, you can move one ball from a box to an adjacent box. Box i
//is adjacent to box j if abs(i - j) == 1. Note that after doing so, there may be
//more than one ball in some boxes.
//
// Return an array answer of size n, where answer[i] is the minimum number of
//operations needed to move all the balls to the iᵗʰ box.
//
// Each answer[i] is calculated considering the initial state of the boxes.
//
//
// Example 1:
//
//
//Input: boxes = "110"
//Output: [1,1,3]
//Explanation: The answer for each box is as follows:
//1) First box: you will have to move one ball from the second box to the first
//box in one operation.
//2) Second box: you will have to move one ball from the first box to the
//second box in one operation.
//3) Third box: you will have to move one ball from the first box to the third
//box in two operations, and move one ball from the second box to the third box in
//one operation.
//
//
// Example 2:
//
//
//Input: boxes = "001011"
//Output: [11,8,5,4,3,4]
//
//
// Constraints:
//
//
// n == boxes.length
// 1 <= n <= 2000
// boxes[i] is either '0' or '1'.
//
//
// Related Topics Array String 👍 2155 👎 84

// leetcode submit region begin(Prohibit modification and deletion)
func minOperations(boxes string) []int {
	result := make([]int, len(boxes))
	//boxes to map. Key is index and value is result
	for i, v := range boxes {
		if v == '1' {
			// add all result to :i
			for j := 0; j < len(boxes); j++ {
				offset := i - j
				if offset < 0 {
					offset = -offset
				}
				result[j] += offset
			}
		}
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)

func Test_MinimumNumberOfOperationsToMoveAllBallsToEachBox(t *testing.T) {

}

func Test_MinimumNumberOfOperationsToMoveAllBallsToEachBox_table(t *testing.T) {
	type args struct {
		boxes string
	}

	tests := []struct {
		name     string
		args     args
		expected []int
	}{
		{
			name: "1",
			args: args{
				boxes: "110",
			},
			expected: []int{1, 1, 3},
		},
		{
			name: "2",
			args: args{
				boxes: "001011",
			},
			expected: []int{11, 8, 5, 4, 3, 4},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := minOperations(test.args.boxes)
			assert.ElementsMatchf(t, test.expected, got, "not match, expected %v, got %v", test.expected, got)
		})
	}

}
